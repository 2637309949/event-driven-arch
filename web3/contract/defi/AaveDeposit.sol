// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/*
AaveDeposit.sol
- 用户先 approve 本合约，allowance 给本合约（ERC20）
- 用户调用 depositToAave(asset, amount)
  1) 合约从用户 transferFrom -> 合约
  2) 合约 approve Pool 合约
  3) 合约 调用 pool.supply(asset, amount, onBehalfOf=user, referralCode=0)
  -> 在 Aave，上述操作将把 aToken credit 给 onBehalfOf（用户）
- withdrawFromAave(asset, amount, to) 由合约调用 pool.withdraw，把资产从 Aave 提回到 to
*/

import "../external/aave-v3-core/contracts/interfaces/IPoolAddressesProvider.sol";
import "../external/aave-v3-core/contracts/interfaces/IPool.sol";
import "../external/openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../external/openzeppelin/contracts/security/ReentrancyGuard.sol";

contract AaveDeposit is ReentrancyGuard {
    IPoolAddressesProvider public immutable provider;

    event DepositedToAave(address indexed user, address indexed asset, uint256 amount);
    event WithdrawnFromAave(address indexed caller, address indexed asset, uint256 amount, address indexed to);

    constructor(address _poolAddressesProvider) {
        require(_poolAddressesProvider != address(0), "provider zero");
        provider = IPoolAddressesProvider(_poolAddressesProvider);
    }

    // internal helper to fetch current Pool
    function _getPool() internal view returns (IPool) {
        address poolAddr = provider.getPool();
        require(poolAddr != address(0), "pool not set");
        return IPool(poolAddr);
    }

    /**
     * @notice 用户把 ERC20 转入本合约并存入 Aave，aToken credit 给 onBehalfOf (= msg.sender)
     * @param asset ERC20 代币地址 (例如 DAI)
     * @param amount 要存入的数量（token decimals）
     *
     * Usage:
     * 1) ERC20.approve(AaveDepositAddress, amount)
     * 2) AaveDeposit.depositToAave(asset, amount)
     */
    function depositToAave(address asset, uint256 amount) external nonReentrant {
        require(amount > 0, "zero amount");
        require(asset != address(0), "zero asset");

        IERC20 token = IERC20(asset);

        // transfer tokens from user to this contract
        bool ok = token.transferFrom(msg.sender, address(this), amount);
        require(ok, "transferFrom failed");

        // get Pool and approve it to pull tokens
        IPool pool = _getPool();

        // reset/approve pattern
        require(token.approve(address(pool), 0), "approve reset failed");
        require(token.approve(address(pool), amount), "approve failed");

        // call supply: this contract supplies tokens, but we credit onBehalfOf = msg.sender
        pool.supply(asset, amount, msg.sender, 0);

        emit DepositedToAave(msg.sender, asset, amount);
    }

    /**
     * @notice 从 Aave 提现，调用 pool.withdraw，将资产提回到 to
     * @param asset ERC20 代币地址
     * @param amount 要提现的数量（可为 type(uint256).max 表示全部）
     * @param to 提款接收地址
     *
     * 注意：谁能调用取决于你合约的业务逻辑（示例中任何人可触发，本示例不做权限控制）
     * 真实环境请严格控制权限（只有 owner / 用户本人才可提取）
     */
    function withdrawFromAave(address asset, uint256 amount, address to) external nonReentrant returns (uint256) {
        require(to != address(0), "zero to");
        IPool pool = _getPool();

        // pool.withdraw 把对应的 underlying token 从 Aave 提回并发送给 to
        uint256 withdrawn = pool.withdraw(asset, amount, to);

        emit WithdrawnFromAave(msg.sender, asset, withdrawn, to);
        return withdrawn;
    }

    // Helper: 允许合约的持有者（或任何）检索合约中误发的 ERC20（仅示例）
    function rescueERC20(address tokenAddress, uint256 amount, address to) external nonReentrant {
        require(to != address(0), "zero to");
        IERC20(tokenAddress).transfer(to, amount);
    }
}
