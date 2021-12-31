import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor, HttpResponse, HttpErrorResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.service';
import { tap } from 'rxjs/operators';

@Injectable()
export class JwtInterceptor implements HttpInterceptor {

    constructor(public auth: AuthService) {}

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
      
        return next.handle(req)
        .pipe(
            tap(
              event => this.handleResponse(req, event),
              error => this.handleError(req, error)
            )
          );
    }

    handleResponse(req: HttpRequest<any>, event: HttpEvent<any>) {
      
        if (event instanceof HttpResponse) {
          // do something if you want
        }
    }
    
    handleError(req: HttpRequest<any>, err: any) {

    if (err.status === 401) {
        this.auth.collectFailedRequest(req);
        // redirect to the login route
        // or show a modal
        }

        console.error('Request for ', req.url,
                ' Response Status ', err.status,
                ' With error ', err.error);
    }
}