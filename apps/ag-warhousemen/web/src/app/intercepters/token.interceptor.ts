import { Injectable } from '@angular/core';
import {
   HttpRequest,
   HttpHandler,
   HttpEvent,
   HttpInterceptor,
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.service';
import { INTERCEPTOR_NO_AUTH_HEADER } from '../services/constants.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {
   constructor(public auth: AuthService) {}

   intercept(
      request: HttpRequest<any>,
      next: HttpHandler
   ): Observable<HttpEvent<any>> {
      // If the incoming request has the special INTERCEPTOR_NO_AUTH_HEADER,
      // it means that we don't have to add the JWT token
      if (request.headers.get(INTERCEPTOR_NO_AUTH_HEADER)) {
         // So just send to the next handler the same request
         // removing the special header INTERCEPTOR_NO_AUTH_HEADER
         return next.handle(
            request.clone({
               headers: request.headers.delete(INTERCEPTOR_NO_AUTH_HEADER),
            })
         );
      }

      // has access token
      if (!!this.auth.getAccessToken()) {
         return next.handle(this.addTokenToRequest(request));
      }

      // If the user is not logged in, just do nothing.
      // We don't have any JWT token to attach to the request
      return next.handle(request);
   }

   // It just clones the request, adding the JWT token to the request
   private addTokenToRequest(request: HttpRequest<any>): HttpRequest<any> {
      return request.clone({
         setHeaders: {
            Authorization: `Bearer ${this.auth.getAccessToken()}`,
         },
      });
   }
}
