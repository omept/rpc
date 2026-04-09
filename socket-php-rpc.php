<?php

function callRPC($method, $params, $id = 1)
{
    $errno = null;
    $errstr = null;
    $fp = fsockopen("127.0.0.1", 1234, $errno, $errstr, 5);

    if (!$fp) {
        throw new Exception("Connection failed: $errstr ($errno)");
    }

    $request = json_encode([
        "method" => $method,
        "params" => [$params],
        "id" => $id
    ]);

    fwrite($fp, $request);

    $response = '';
    while (true) {
        $chunk = fread($fp, 1024);
        if ($chunk === false || $chunk === '') {
            break;
        }

        $response .= $chunk;

        $decoded = json_decode($response, true);
        if ($decoded !== null) {
            break;
        }
    }

    fclose($fp);

    $decoded = json_decode($response, true);
    if ($decoded == null) {
        throw new Exception("decoded content is empty");
    }

    if ((isset($decoded['error']) && $decoded['error'] !== null)) {
        throw new Exception($decoded['error']);
    }

    return $decoded['result'];
}


try {
    $result = callRPC("Arith.Multiply", ["A" => 17, "B" => 8]);
    echo "Multiply Result: " . $result . PHP_EOL;

    $division = callRPC("Arith.Divide", ["A" => 17, "B" => 8]);
    echo "Divide Result: ";
    print_r($division);
} catch (Exception $e) {
    echo "RPC Error: " . $e->getMessage();
}
