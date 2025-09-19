solc --bin --abi PracticalNFT.sol -o . --base-path . --include-path .. --allow-paths .. --overwrite
abigen --bin=PracticalNFT.bin --abi=PracticalNFT.abi --pkg=nft --out=PracticalNFT.go