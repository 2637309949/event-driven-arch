// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
 NFT contract using OpenZeppelin v4.8 (raw imports suitable for Remix).
 Features:
  - ERC721Enumerable
  - Owner controls: mint, toggle sale/whitelist, set base URI, withdraw
  - Public mint + Merkle-based whitelist mint
  - Reveal support (unrevealedURI)
  - Basic anti-bot/anti-reentrancy checks
 Note: 在生产部署前请审计并根据需要开启/加入版税 (ERC2981)、gas 优化等。
*/

import "../openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "../openzeppelin/contracts/access/Ownable.sol";
import "../openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "../openzeppelin/contracts/utils/Strings.sol";

contract NFT is ERC721Enumerable, Ownable, ReentrancyGuard {
    using Strings for uint256;

    // Sale params
    uint256 public maxSupply;
    uint256 public mintPrice;
    uint256 public maxPerTx;
    uint256 public maxPerWallet;

    // State flags
    bool public saleActive = false;
    bool public whitelistActive = false;
    bool public revealed = false;

    // URIs
    string private baseTokenURI;
    string public unrevealedURI;

    // Merkle root for whitelist
    bytes32 public merkleRoot;

    // Track per-wallet minted count (for maxPerWallet enforcement)
    mapping(address => uint256) public mintedPerWallet;

    // Events
    event Minted(address indexed minter, uint256 indexed tokenId);
    event BaseURISet(string baseURI);
    event UnrevealedURISet(string unrevealedURI);
    event SaleToggled(bool active);
    event WhitelistToggled(bool active);
    event Revealed(bool revealed);

    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _maxSupply,
        uint256 _mintPrice,
        uint256 _maxPerTx,
        uint256 _maxPerWallet,
        string memory _unrevealedURI
    ) ERC721(_name, _symbol) {
        maxSupply = _maxSupply;
        mintPrice = _mintPrice;
        maxPerTx = _maxPerTx;
        maxPerWallet = _maxPerWallet;
        unrevealedURI = _unrevealedURI;
    }

    // ========== Modifiers ==========
    modifier canMint(uint256 quantity) {
        require(quantity > 0, "quantity>0");
        require(quantity <= maxPerTx, "exceed maxPerTx");
        require(totalSupply() + quantity <= maxSupply, "exceed maxSupply");
        _;
    }

    // ========== Minting ==========
    function publicMint(uint256 quantity) external payable nonReentrant canMint(quantity) {
        require(saleActive, "sale not active");
        require(!whitelistActive, "public mint closed (whitelist active)");
        require(msg.value >= mintPrice * quantity, "insufficient funds");
        require(mintedPerWallet[msg.sender] + quantity <= maxPerWallet, "exceed wallet limit");

        mintedPerWallet[msg.sender] += quantity;
        _batchMint(msg.sender, quantity);
    }

    function whitelistMint(uint256 quantity, bytes32[] calldata proof) external payable nonReentrant canMint(quantity) {
        require(whitelistActive, "whitelist not active");
        require(msg.value >= mintPrice * quantity, "insufficient funds");
        require(mintedPerWallet[msg.sender] + quantity <= maxPerWallet, "exceed wallet limit");
        // verify merkle proof (leaf as keccak256(abi.encodePacked(address)))
        bytes32 leaf = keccak256(abi.encodePacked(msg.sender));
        require(MerkleProof.verify(proof, merkleRoot, leaf), "invalid proof");

        mintedPerWallet[msg.sender] += quantity;
        _batchMint(msg.sender, quantity);
    }

    // Owner-only mint (for team / airdrop)
    function ownerMint(address to, uint256 quantity) external onlyOwner canMint(quantity) {
        _batchMint(to, quantity);
    }

    function _batchMint(address to, uint256 quantity) internal {
        uint256 start = totalSupply();
        for (uint256 i = 0; i < quantity; i++) {
            uint256 tokenId = start + i + 1; // tokenId starts from 1
            _safeMint(to, tokenId);
            emit Minted(to, tokenId);
        }
    }

    // ========== Metadata ==========
    function setBaseURI(string memory _baseURI) external onlyOwner {
        baseTokenURI = _baseURI;
        emit BaseURISet(_baseURI);
    }

    function setUnrevealedURI(string memory _unrevealedURI) external onlyOwner {
        unrevealedURI = _unrevealedURI;
        emit UnrevealedURISet(_unrevealedURI);
    }

    function reveal() external onlyOwner {
        revealed = true;
        emit Revealed(true);
    }

    function hide() external onlyOwner {
        revealed = false;
        emit Revealed(false);
    }

    function tokenURI(uint256 tokenId) public view virtual override returns (string memory) {
        require(_exists(tokenId), "nonexistent token");
        if (!revealed) {
            return unrevealedURI;
        }
        string memory base = baseTokenURI;
        return bytes(base).length > 0 ? string(abi.encodePacked(base, tokenId.toString())) : "";
    }

    // ========== Admin controls ==========
    function toggleSale() external onlyOwner {
        saleActive = !saleActive;
        emit SaleToggled(saleActive);
    }

    function toggleWhitelist() external onlyOwner {
        whitelistActive = !whitelistActive;
        emit WhitelistToggled(whitelistActive);
    }

    function setMerkleRoot(bytes32 _root) external onlyOwner {
        merkleRoot = _root;
    }

    function setMintPrice(uint256 _mintPrice) external onlyOwner {
        mintPrice = _mintPrice;
    }

    function setMaxPerTx(uint256 _maxPerTx) external onlyOwner {
        maxPerTx = _maxPerTx;
    }

    function setMaxPerWallet(uint256 _maxPerWallet) external onlyOwner {
        maxPerWallet = _maxPerWallet;
    }

    function setMaxSupply(uint256 _maxSupply) external onlyOwner {
        require(_maxSupply >= totalSupply(), "less than minted");
        maxSupply = _maxSupply;
    }

    // ========== Withdraw ==========
    function withdraw(address payable to) external onlyOwner nonReentrant {
        uint256 bal = address(this).balance;
        require(bal > 0, "no balance");
        // Example: simple withdraw entire balance to `to`.
        // For split payments, implement PaymentSplitter or custom split logic here.
        (bool ok, ) = to.call{value: bal}("");
        require(ok, "withdraw failed");
    }

    // ========== Overrides ==========
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721Enumerable) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
