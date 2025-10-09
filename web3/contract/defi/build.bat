solc --bin --abi AaveDeposit.sol -o . --base-path . --include-path .. --allow-paths .. --overwrite
abigen --bin=AaveDeposit.bin --abi=AaveDeposit.abi --pkg=defi --out=AaveDeposit.go