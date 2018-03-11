import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { StartLoginUser, ContLoginUser } from './login.component';
import { TestURL, APIURL } from '../api/api.constants';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    // 'Authorization': 'my-auth-token' --------------- add here
  })
};

const httpOptionsAuthorized = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    'Authorization': 'Bearer ' + this.token
  })
};

@Injectable()
export class LoginTestService {
  public token: string;
  public loginState: string;
  private handleError: HandleError;
  private email: string;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('LoginTestService');

    const currentUser = JSON.parse(localStorage.getItem('currentUser'));
    this.token = currentUser && currentUser.token;
  }

  startLoginUser (login: StartLoginUser): Observable<StartLoginUser> {
    return this.http.post<StartLoginUser>(TestURL + 'login-start', login, httpOptions)
      .pipe(
        catchError(this.handleError('loginUser', login))
      );
  }

  contLoginUser (login: ContLoginUser): Observable<boolean> {
    return this.http.post<ContLoginUser>(TestURL + 'login-finish', login, httpOptions).map(
      response => {
        if (response.token) {
          this.token = response.token;
          this.loginState = response.loginState;
          localStorage.setItem('currentUser', JSON.stringify({
                                id: response.id,
                                email: login.email,
                                loginState: this.loginState,
                                token: this.token }));
          return true;
        }
        return false;
    });
  }

  tryPing (): Observable<boolean> {
    return this.http.get(TestURL + 'ping', httpOptionsAuthorized).map(
      response => {
        console.log(response);
          return true;
      });
  }

  logout(): void {
    // clear token remove user from local storage to log user out
    this.token = null;
    localStorage.removeItem('currentUser');
  }
}
