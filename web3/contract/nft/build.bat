solc --bin --abi NFT.sol -o . --base-path . --include-path .. --allow-paths .. --overwrite
abigen --bin=NFT.bin --abi=NFT.abi --pkg=nft --out=NFT.go