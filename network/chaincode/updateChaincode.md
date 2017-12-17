//enter container
$ docker exec -ti cli bash

//install the new chaincode ,it needs modify -n or -v
$ peer chaincode install -n mycc -v 2.0 -p github.com/hyperledger/fabric/examples/chaincode/go/ 

//upgrade
$ peer chaincode upgrade -n mycc -v 2.0 -c '{"Args":[""]}' -C mychannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

//exit
exit