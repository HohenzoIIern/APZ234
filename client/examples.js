// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require('./server/client');

const client = channels.Client('http://localhost:8080');

const addNameServer = 'qw';
const addCPUServer = 2;
const addSpaceServer = 1;

// Scenario 1: Display available channels.
client.listServer()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available server:');
        list.forEach((c) => console.log(c.name, c.cpu_count, c.totalDiskSpace));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

// Scenario 1: Display available channels.
client.listDisk()
    .then((list) => {
        console.log('=== Scenario 2 ===');
        console.log('Available disk:');
        list.forEach((c) => console.log(c.id, c.space, c.server_id));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

// Scenario 1: Display available channels.
client.createDisk(addSpaceServer)
    .then((resp) => {
        console.log('=== Scenario 3 ===');
        console.log('Create disk response:', resp);
        return resp
    }).then((disk) => {
        return client.createServer(addNameServer, addCPUServer)
            .then((resp) => {
                console.log('=== Scenario 4 ===');
                console.log('Create server response:', resp);
                return resp
            }).then((server) => {
                return client.addDiskToServer(server.id, disk.id)
            })
            .then((resp) => {
                console.log('=== Scenario 5 ===');
                console.log('Add disk to server response:', resp);
            })
            .catch((e) => {
                console.log(`Problem listing available channels: ${e.message}`);
            });
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });