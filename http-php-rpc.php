<?php

function callRPC($method, $params, $id = 1)
{
    $payload = json_encode([
        "method" => $method,
        "params" => [$params],
        "id" => $id
    ]);

    $ch = curl_init("http://127.0.0.1:1234/rpc");

    curl_setopt_array($ch, [
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_POST => true,
        CURLOPT_HTTPHEADER => ['Content-Type: application/json'],
        CURLOPT_POSTFIELDS => $payload,
    ]);

    $response = curl_exec($ch);

    if ($response === false) {
        throw new Exception(curl_error($ch));
    }

    unset($ch);

    $decoded = json_decode($response, true);

    if (isset($decoded['error']) && $decoded['error']) {
        throw new Exception($decoded['error']);
    }

    return $decoded['result'];
}


try {
    echo "Multiply: " . callRPC("Arith.Multiply", ["A" => 6, "B" => 3]) . PHP_EOL;

    print_r(callRPC("Arith.Divide", ["A" => 10, "B" => 3]));
} catch (Exception $e) {
    echo "Error: " . $e->getMessage();
}
