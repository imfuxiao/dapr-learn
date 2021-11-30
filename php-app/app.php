<?php

$method = $_SERVER['REQUEST_METHOD'];
$uri = $_SERVER['REQUEST_URI'];

$headers = array();
foreach ($_SERVER as $key => $value) {
    if (strpos($key, 'HTTP_') == 0 && strlen($key) >5) {
        $header = str_replace(' ', '-', ucwords(str_replace('_', ' ', strtolower(substr($key, 5)))));
        $headers[$header] = $value;
    }
}

echo json_encode(array('method'=>$method, 'uri' => $uri, 'headers' => $headers));

?>