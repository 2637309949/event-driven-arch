solc --abi --bin Store.sol -o ./ --overwrite 
abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go


Get-Process -Id (Get-NetTCPConnection -LocalPort 8080).OwningProcess | Stop-Process -Force
