import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { TestUser, ContTestUser } from './login-test.component';
import { TestURL, APIURL } from '../api/api.constants';
import { UserConstantsService } from '../dashboard/user-constants/user-constants.service';

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
  public testLevel: number;
  private handleError: HandleError;
  private email: string;

  constructor(private http: HttpClient,
              httpErrorHandler: HttpErrorHandler,
              private userConstants: UserConstantsService) {
    this.handleError = httpErrorHandler.createHandleError('LoginTestService');

    const currentUser = JSON.parse(localStorage.getItem('currentUser'));
    this.token = currentUser && currentUser.token;
  }

  startLoginUser (login: TestUser): Observable<TestUser> {
    return this.http.post<TestUser>(APIURL + 'test/login-start', login, httpOptions);
      // .pipe(
      //   catchError(this.handleError('loginUser', login))
      // );
  }

  contLoginUser (login: ContTestUser): Observable<boolean> {
    return this.http.post<ContTestUser>(APIURL + 'test/login-finish', login, httpOptions).map(
      response => {
        if (response.token) {
          this.userConstants.ID = response.id;
          this.userConstants.Email = response.email;
          this.userConstants.TestLevel = response.testLevel;
          this.userConstants.Token = response.token;
          localStorage.setItem('currentUser', JSON.stringify({
                                id: response.id,
                                email: login.email,
                                testLevel: login.testLevel,
                                init: response.init,
                                token: response.token }));
          return true;
        }
        return false;
    });
  }

  logout(): void {
    // clear token remove user from local storage to log user out
    this.token = null;
    localStorage.removeItem('currentUser');
  }
}
