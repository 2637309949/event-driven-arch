solc --bin --abi SimpleSwap.sol -o . --base-path . --include-path .. --allow-paths .. --overwrite
abigen --bin=SimpleSwap.bin --abi=SimpleSwap.abi --pkg=swap --out=SimpleSwap.go