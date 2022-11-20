import { Controller, Logger } from '@nestjs/common';
import { EventPattern, MessagePattern } from '@nestjs/microservices';

@Controller()
export class AppController {
  private logger = new Logger(AppController.name);

  @EventPattern()
  getGreetingMessage(name: string): string {
    this.logger.debug('Received message from another service::: ' + name);
    return `Hello ${name}`;
  }
}
