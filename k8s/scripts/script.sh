#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Build your first network (BYFN) end-to-end test"
echo
CHANNEL_NAME="mychannel"
DELAY="3"
: ${CHANNEL_NAME:="mychannel"}
: ${TIMEOUT:="60"}
COUNTER=1
MAX_RETRY=5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/fabric-net.svc.cluster.local/orderers/orderer.fabric-net.svc.cluster.local/msp/tlscacerts/tlsca.fabric-net.svc.cluster.local-cert.pem

echo "Channel name : "$CHANNEL_NAME

# verify the result of the end-to-end test
verifyResult () {
  if [ $1 -ne 0 ] ; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
    echo
      exit 1
  fi
}

setGlobals () {

  if [ $1 -eq 0 ] ; then
    CORE_PEER_LOCALMSPID="Org1MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.fabric-net.svc.cluster.local/peers/peer0-org1.fabric-net.svc.cluster.local/tls/ca.crt
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.fabric-net.svc.cluster.local/users/Admin@org1.fabric-net.svc.cluster.local/msp
    CORE_PEER_ADDRESS=peer0-org1.fabric-net.svc.cluster.local:7051
  elif [ $1 -eq 1 ]; then
    CORE_PEER_LOCALMSPID="Org2MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.fabric-net.svc.cluster.local/peers/peer0-org2.fabric-net.svc.cluster.local/tls/ca.crt
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.fabric-net.svc.cluster.local/users/Admin@org2.fabric-net.svc.cluster.local/msp
    CORE_PEER_ADDRESS=peer0-org2.fabric-net.svc.cluster.local:7051
  elif [ $1 -eq 2 ]; then
    CORE_PEER_LOCALMSPID="Org3MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.fabric-net.svc.cluster.local/peers/peer0-org3.fabric-net.svc.cluster.local/tls/ca.crt
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.fabric-net.svc.cluster.local/users/Admin@org3.fabric-net.svc.cluster.local/msp
    CORE_PEER_ADDRESS=peer0-org3.fabric-net.svc.cluster.local:7051
  elif [ $1 -eq 3 ]; then
    CORE_PEER_LOCALMSPID="Org4MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.fabric-net.svc.cluster.local/peers/peer0-org4.fabric-net.svc.cluster.local/tls/ca.crt
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.fabric-net.svc.cluster.local/users/Admin@org4.fabric-net.svc.cluster.local/msp
    CORE_PEER_ADDRESS=peer0-org4.fabric-net.svc.cluster.local:7051
  fi

  env |grep CORE
}

createChannel() {
  setGlobals 0

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    peer channel create -o orderer.fabric-net.svc.cluster.local:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
  else
    peer channel create -o orderer.fabric-net.svc.cluster.local:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
  fi
  res=$?
  cat log.txt
  verifyResult $res "Channel creation failed"
  echo "===================== Channel \"$CHANNEL_NAME\" is created successfully ===================== "
  echo
}

updateAnchorPeers() {
  PEER=$1
  setGlobals $PEER

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    peer channel update -o orderer.fabric-net.svc.cluster.local:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx >&log.txt
  else
    peer channel update -o orderer.fabric-net.svc.cluster.local:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
  fi
  res=$?
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers for org \"$CORE_PEER_LOCALMSPID\" on \"$CHANNEL_NAME\" is updated successfully ===================== "
  sleep $DELAY
  echo
}

## Sometimes Join takes time hence RETRY atleast for 5 times
joinWithRetry () {
  peer channel join -b $CHANNEL_NAME.block  >&log.txt
  res=$?
  cat log.txt
  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=` expr $COUNTER + 1`
    echo "PEER$1 failed to join the channel, Retry after 2 seconds"
    sleep $DELAY
    joinWithRetry $1
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, PEER$ch has failed to Join the Channel"
}

joinChannel () {
  for ch in 0 1 2 3; do
    setGlobals $ch
    joinWithRetry $ch
    echo "===================== PEER$ch joined on the channel \"$CHANNEL_NAME\" ===================== "
    sleep $DELAY
    echo
  done
}

installChaincode () {
  PEER=$1
  setGlobals $PEER
  peer chaincode install -n mycc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go >&log.txt
  res=$?
  cat log.txt
        verifyResult $res "Chaincode installation on remote peer PEER$PEER has Failed"
  echo "===================== Chaincode is installed on remote peer PEER$PEER ===================== "
  echo
}

instantiateChaincode () {
  PEER=$1
  setGlobals $PEER
  # while 'peer chaincode' command can get the orderer endpoint from the peer (if join was successful),
  # lets supply it directly as we know it using the "-o" option
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    peer chaincode instantiate -o orderer.fabric-net.svc.cluster.local:7050 -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')" >&log.txt
  else
    peer chaincode instantiate -o orderer.fabric-net.svc.cluster.local:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')" >&log.txt
  fi
  res=$?
  cat log.txt
  verifyResult $res "Chaincode instantiation on PEER$PEER on channel '$CHANNEL_NAME' failed"
  echo "===================== Chaincode Instantiation on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
  echo
}

## Create channel
echo "Creating channel..."
createChannel

## Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

## Set the anchor peers for each org in the channel
echo "Updating anchor peers for org1..."
updateAnchorPeers 0
echo "Updating anchor peers for org2..."
updateAnchorPeers 1
echo "Updating anchor peers for org3..."
updateAnchorPeers 2
echo "Updating anchor peers for org4..."
updateAnchorPeers 3

## Install chaincode on Peer0/Org1, Peer0/Org2 and Peer0/Org3
echo "Install chaincode on org1/peer0..."
installChaincode 0
echo "Install chaincode on org2/peer0..."
installChaincode 1
echo "Install chaincode on org3/peer0..."
installChaincode 2
echo "Install chaincode on org4/peer0..."
installChaincode 3

#Instantiate chaincode on Peer0/Org3
echo "Instantiating chaincode on org4/peer0..."
instantiateChaincode 3


echo
echo "========= All GOOD, BYFN execution completed =========== "
echo

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0

