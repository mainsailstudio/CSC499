import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

// import { StartRegisterUser } from './register.component';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    // 'Authorization': 'my-auth-token' --------------- add here
  })
};

@Injectable()
export class LoginUserService {
  registerUrl = 'http://13.92.156.114:8080/login';  // URL to web api
  private handleError: HandleError;
  private email: string;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('LoginUserService');
  }

  loginUser (email: string) {
    return this.http.post(this.registerUrl, email, httpOptions)
      .pipe(
        catchError(this.handleError('registerUser', email))
      );
  }

}
