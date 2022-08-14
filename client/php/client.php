<?php
require dirname(__FILE__).'/vendor/autoload.php';
function login($name, $pwd, $hostname)
{

    // 初始化一个客户端实例
    $client = new \Login\loginClient($hostname, [
        'credentials' => Grpc\ChannelCredentials::createInsecure(),
    ]);
    // 初始化一个请求类
    $request = new \Login\Request();
    // 参数赋值
    $st = time();
    $orgStr = sprintf("name=%name&pwd=%pwd&stime=%st&key=123", $name, $pwd, $st);
    $sign = md5($orgStr);

    $request->setName($name);
    $request->setPwd($pwd);
    $request->setStime($st);
    $request->setSign($sign);
    // 请求服务
    list($response, $status) = $client->Login($request)->wait();
    // 响应处理
    if ($status->code !== Grpc\STATUS_OK) {
        echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
        exit(1);
    }
    echo $response->getMessage() . PHP_EOL;
}



$name = 'afei';
$pwd = md5('123456');
$hostname = 'localhost:8322';

login($name, $pwd, $hostname);
