<?php

namespace App\Console\Commands;

use App\Service\Amqp\MessageHandlerService;
use Illuminate\Console\Command;
use PhpAmqpLib\Connection\AMQPStreamConnection;
class AmqpListenCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'app:amqp-listen';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        try {
            $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
            $channel = $connection->channel();

            $queue_name = 'course_queue';
            $channel->queue_declare($queue_name, false, true, false, false);

            $messageHandler = new MessageHandlerService();
            $callback = function ($msg) use ($messageHandler) {
                $type = $msg->get('type');
                $payload = json_decode($msg->getBody(), true);
                $messageHandler->handleMessage($type, $payload);
            };

            $channel->basic_consume($queue_name, '', false, true, false, false, $callback);

            while ($channel->is_consuming()) {
                $channel->wait();
            }

        } catch (\Exception $exception) {
            $this->error($exception->getMessage());
        }
    }
}
