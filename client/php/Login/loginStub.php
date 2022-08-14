<?php
namespace Login;


class LoginStub {


    public function Login123(
        \Login\Request $request,
        \Grpc\ServerContext $context
    ): ?\Login\LoginReply {
        $context->setStatus(\Grpc\Status::unimplemented());
        return null;
    }

    /**
     * Get the method descriptors of the service for server registration
     *
     * @return array of \Grpc\MethodDescriptor for the service methods
     */
    public final function getMethodDescriptors(): array
    {
        return [
            '/Login/Login' => new \Grpc\MethodDescriptor(
                $this,
                'Login',
                '\Login\Request',
                \Grpc\MethodDescriptor::UNARY_CALL
            ),
        ];
    }

}
