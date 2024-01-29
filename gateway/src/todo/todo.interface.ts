import { Observable } from 'rxjs';
import { CreateTodoDTO, TodoByIdDTO, UpdateTodoContentDTO } from './todo.dto';

export interface Todo {
  id: number;
  content: string;
  isComplete: boolean;
}

export interface Todos {
  todos: Todo[];
}

export interface TodoGRPCServive {
  getTodo: (data: TodoByIdDTO) => Observable<Todo>;
  getAllTodos: (data: Record<string, never>) => Observable<Todos>;
  createTodo: (data: CreateTodoDTO) => Observable<Todo>;
  updateTodoContent: (data: UpdateTodoContentDTO) => Observable<Todo>;
  markTodoAsComplete: (data: TodoByIdDTO) => Observable<Todo>;
  markTodoAsInComplete: (data: TodoByIdDTO) => Observable<Todo>;
  deleteTodo: (data: TodoByIdDTO) => Observable<Todo>;
}
