// --------------------      Reference   --------------------- //
//  https://blog.mattianatali.dev/jwt-authentication-angular/  //
// ----------------------------------------------------------- //
import { Injectable } from '@angular/core';
import {
   HttpRequest,
   HttpHandler,
   HttpEvent,
   HttpInterceptor,
   HttpErrorResponse,
} from '@angular/common/http';
import { catchError, mergeMap, Observable, tap } from 'rxjs';
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

      // If the user is logged in
      if (!!this.auth.getAccessToken()) {
         return (
            next
               // Then add the token and continue the API call.
               .handle(this.addTokenToRequest(request))
               .pipe(
                  // But if the call has failed, try to recover it using retryWhenExpiredToken method
                  catchError(this.retryWhenExpiredToken(request, next)),
                  tap({
                     error: (err: HttpErrorResponse) => {
                        //  If the call, after our refresh method, is still failing with
                        //  401 UNAUTHORIZED status code
                        if (err.status === 401) {
                           // Then logout the user, nothing else we can do.
                           this.auth.logout();
                        }
                     },
                  })
               )
         );
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

   private retryWhenExpiredToken(
      req: HttpRequest<any>,
      next: HttpHandler
   ): (response: any) => Observable<any> {
      return (response) => {
         const refreshToken = this.auth.getRefreshToken();
         //  If the status code is 401 UNAUTHORIZED and we don't have the refresh token,
         //  or the status code is different from 401 UNAUTHORIZED
         if (
            (response.status === 401 && !refreshToken) ||
            response.status !== 401
         ) {
            // We can do nothing, just re-raise the error and continue.
            throw response;
         }

         // Otherwise, call the backend to refresh our token
         return this.auth.refresh().pipe(
            // Then re-call the failed request with the brand new token
            mergeMap(() => {
               return next.handle(this.addTokenToRequest(req));
            })
         );
      };
   }
}
