import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ErrorMsgService {

  errors = new Map<number, string>([
    [0, 'Action is sucessfull' ],
    [1, 'Login failed. Please check your user id and password'],
    [2, 'User already exist. Please use "Forget Password" option to reset your password'],
    [3, 'User does not eixst. Please check user id and try again'],
    [4, 'User is not authorized'],
    [5, 'Invalid request, please try after some time'],
    [6, 'Registration failed because server has issue, please try after some time'],
  ]);

  constructor() { }

  public get(obj: ErrorObject) : string {

    if (this.errors.has(obj.Code)) {
      return this.errors.get(obj.Code) as string;
    }

    if (!!obj.CodeDesc) {
      return obj.CodeDesc
    }

    if (!!obj.Desc) {
      return obj.Desc
    }
    
    return 'Unknown error, please try after some time';
  }
}

export interface ErrorObject {
  Code: number;
  CodeDesc: string;
  Desc: string;
}