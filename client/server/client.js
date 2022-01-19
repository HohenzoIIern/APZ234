const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listServer: () => client.get('/server'),
        createServer: (name, cpu_count) => client.post('/server', { name, cpu_count }),
        addDiskToServer: (server_id, disk_id) => client.patch('/server', { server_id, disk_id }),
        listDisk: () => client.get('/disk'),
        createDisk: (space) => client.post('/disk', { space }),
    }
};

module.exports = { Client };
