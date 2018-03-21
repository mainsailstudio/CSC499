import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { RegisterTestUser } from './register-test.component';
import { APIURL, TestURL } from '../api/api.constants';
import { UserConstantsService } from '../dashboard/user-constants/user-constants.service';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    // 'Authorization': 'my-auth-token' --------------- add here
  })
};

@Injectable()
export class RegisterTestService {
  private handleError: HandleError;
  public token;
  public testLevel;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler,
    private userConstants: UserConstantsService) {
    this.handleError = httpErrorHandler.createHandleError('RegisterStartService');
  }

  /** POST: add a new user to the database */
  registerTest (register: RegisterTestUser): Observable<RegisterTestUser> {
    return this.http.post<RegisterTestUser>(APIURL + 'test/register', register, httpOptions).map(
      response => {
        // first check if the user is already initialized, ie registered
        if (response.init) {
          return null;
        } else if (response.token) {
          this.userConstants.ID = response.id;
          this.userConstants.Email = response.email;
          console.log('====== User constant email is ======= ' + this.userConstants.Email);
          this.userConstants.TestLevel = response.testLevel;
          this.userConstants.Token = response.token;
          this.userConstants.Init = response.init;
          localStorage.setItem('currentUser', JSON.stringify({
                                id: response.id,
                                email: response.email,
                                testLevel: response.testLevel,
                                init: response.init,
                                token: response.token }));
          return response;
        }
        return response;
    });
  }

  loginTestUser (loginUser: RegisterTestUser): Observable<boolean> {
    return this.http.post<RegisterTestUser>(APIURL + 'test/login-token', loginUser, httpOptions).map(
      response => {
        if (response.token) {
          this.token = response.token;
          this.testLevel = response.testLevel;
          localStorage.setItem('currentUser', JSON.stringify({
                                id: response.id,
                                email: loginUser.email,
                                testLevel: this.testLevel,
                                token: this.token }));
          return true;
        }
        return false;
    });
  }

}
