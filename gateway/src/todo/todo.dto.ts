export interface TodoByIdDTO {
  id: number;
}
export interface CreateTodoDTO {
  content: string;
}
export interface UpdateTodoContentDTO {
  id: number;
  content: string;
}
// export interface MarkTodoAsCompleteDTO {
//   id: number;
// }
// export interface MarkTodoAsInCompleteDTO {
//   id: number;
// }
// export interface DeleteTodoDTO {
//   id: number;
// }
