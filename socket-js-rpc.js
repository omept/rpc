const net = require('net');
const PORT = 1234;
const SERVER = '127.0.0.1';
function callRPC(method, params, id = 1) {
    return new Promise((resolve, reject) => {
        const client = new net.Socket();

        client.connect(PORT, SERVER, () => {
            const request = {
                method: method,
                params: [params],
                id: id
            };

            client.write(JSON.stringify(request));
        });

        let dataBuffer = '';

        client.on('data', (data) => {
            dataBuffer += data.toString();

            try {
                const response = JSON.parse(dataBuffer);
                client.destroy();

                if (response.error) {
                    reject(response.error);
                } else {
                    resolve(response.result);
                }
            } catch (e) {
                reject(e);
            }
        });

        client.on('error', (err) => {
            reject(err);
        });
    });
}


(async () => {
    try {
        const result = await callRPC("Arith.Multiple", { A: 17, B: 8 });
        console.log("Multiply Result:", result);

        const division = await callRPC("Arith.Divide", { A: 17, B: 8 });
        console.log("Divide Result:", division);
    } catch (err) {
        console.error("RPC Error:", err);
    }
})();