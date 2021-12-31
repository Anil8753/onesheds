import { HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  cachedRequests: Array<HttpRequest<any>> = [];

  constructor() { }

  public setToken(token: string) {
    localStorage.setItem('token', token);
  }

  public getToken(): string | null {
    return localStorage.getItem('token');
  }

  // public isAuthenticated(): boolean {
  //   // get the token
  //   const token = this.getToken() as string;
  //   // return a boolean reflecting 
  //   // whether or not the token is expired
  //   return tokenNotExpired(token);
  // }

  public collectFailedRequest(request: HttpRequest<any>): void {
    this.cachedRequests.push(request);
  }
    public retryFailedRequests(): void {
    // retry the requests. this method can
    // be called after the token is refreshed
  }
}
