async function callRPC(method, params, id = 1) {
    const res = await fetch('http://127.0.0.1:1234/rpc', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            method: method,
            params: [params],
            id: id
        })
    });

    const data = await res.json();

    if (data.error) {
        throw new Error(data.error);
    }

    return data.result;
}


(async () => {
    try {
        const mul = await callRPC("Arith.Multiply", { A: 17, B: 8 });
        console.log("Multiply:", mul);

        const div = await callRPC("Arith.Divide", { A: 17, B: 8 });
        console.log("Divide:", div);
    } catch (e) {
        console.error(e);
    }
})();