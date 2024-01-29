import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Patch,
  Delete,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { TodoService } from './todo.service';
import { Todo } from './todo.interface';
import { CreateTodoDTO, TodoByIdDTO, UpdateTodoContentDTO } from './todo.dto';

@Controller('todos')
export class TodoController {
  constructor(private todoService: TodoService) {}

  @Get('/:id')
  getTodo(@Param() params: TodoByIdDTO): Observable<Todo> {
    return this.todoService.getTodo(params);
  }

  @Get('/')
  getAllTodo(): Observable<Todo[]> {
    return this.todoService.getAllTodos();
  }

  @Post('/')
  createTodo(@Body() body: CreateTodoDTO): Observable<Todo> {
    return this.todoService.createTodo(body);
  }

  @Patch('/:id')
  updateTodoContent(
    @Param() params: TodoByIdDTO,
    @Body() body: CreateTodoDTO,
  ): Observable<Todo> {
    const data = { ...params, ...body } as UpdateTodoContentDTO;
    return this.todoService.updateTodoContent(data);
  }

  @Patch('/:id/mark-as-complete')
  markTodoAsComplete(@Param() params: TodoByIdDTO): Observable<Todo> {
    return this.todoService.markTodoAsComplete(params);
  }

  @Patch('/:id/mark-as-incomplete')
  markTodoAsInComplete(@Param() params: TodoByIdDTO): Observable<Todo> {
    return this.todoService.markTodoAsInComplete(params);
  }

  @Delete('/:id')
  deleteTodo(@Param() params: TodoByIdDTO): Observable<Todo> {
    return this.todoService.deleteTodo(params);
  }
}
