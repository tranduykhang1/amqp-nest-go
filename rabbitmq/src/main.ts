import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);

  const queueName = configService.get('RABBITMQ_QUEUE_NAME');
  const rabbitClient = configService.get('RABBITMQ_CLIENT');

  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.RMQ,
    options: {
      urls: [rabbitClient],
      queue: queueName,
      noAck: true,
      queueOptions: {
        durable: true,
      },
    },
  });

  app.startAllMicroservices();
  await app.listen(3000);
}

bootstrap();
