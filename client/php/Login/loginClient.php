<?php
namespace Login;
class loginClient extends \Grpc\BaseStub{

    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }


    public function Login(\Login\Request $var = null,$metadata = [], $options = [])
    {
        return $this->_simpleRequest('/Login/Login',
            $var,
            ['\Login\Response', 'decode'],
            $metadata, $options);
    }
}