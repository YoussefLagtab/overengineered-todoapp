import { Injectable, OnModuleInit } from '@nestjs/common';
import { Client, ClientGrpc, Transport } from '@nestjs/microservices';
import { Observable, map } from 'rxjs';
import { Todo, TodoGRPCServive } from './todo.interface';
import { join } from 'path';
import { env } from 'src/config/env';
import { CreateTodoDTO, TodoByIdDTO, UpdateTodoContentDTO } from './todo.dto';

@Injectable()
export class TodoService implements OnModuleInit {
  @Client({
    transport: Transport.GRPC,
    options: {
      url: env.TODOS_SERVICE_URL,
      package: 'todo',
      protoPath: join(__dirname, '../../../proto/todos.proto'),
    },
  })
  client: ClientGrpc;

  private todoGRPCService: TodoGRPCServive;

  onModuleInit() {
    this.todoGRPCService =
      this.client.getService<TodoGRPCServive>('TodoService');
  }

  getTodo(data: TodoByIdDTO): Observable<Todo> {
    return this.todoGRPCService.getTodo(data);
  }
  getAllTodos(): Observable<Todo[]> {
    return this.todoGRPCService.getAllTodos({}).pipe(map((res) => res.todos));
  }
  createTodo(data: CreateTodoDTO): Observable<Todo> {
    return this.todoGRPCService.createTodo(data);
  }
  updateTodoContent(data: UpdateTodoContentDTO): Observable<Todo> {
    return this.todoGRPCService.updateTodoContent(data);
  }
  markTodoAsComplete(data: TodoByIdDTO): Observable<Todo> {
    return this.todoGRPCService.markTodoAsComplete(data);
  }
  markTodoAsInComplete(data: TodoByIdDTO): Observable<Todo> {
    return this.todoGRPCService.markTodoAsInComplete(data);
  }
  deleteTodo(data: TodoByIdDTO): Observable<Todo> {
    return this.todoGRPCService.deleteTodo(data);
  }
}
