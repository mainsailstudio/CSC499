import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { RegisterTestUser } from '../register/register-test.component';
import { APIURL, TestURL } from '../api/api.constants';
import { UserConstantsService } from '../dashboard/user-constants/user-constants.service';

const httpOptions = {
    headers: new HttpHeaders({
      'Content-Type':  'application/json',
      // 'Authorization': 'my-auth-token' --------------- add here
    })
  };

@Injectable()
export class GetTokenService {
  private handleError: HandleError;
  public token;
  public testLevel;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler,
    private userConstants: UserConstantsService) {
    this.handleError = httpErrorHandler.createHandleError('FussFreeTokenBaby');
  }

  /** get a fuss free token */
  getFussFreeToken (register: RegisterTestUser): Observable<RegisterTestUser> {
    return this.http.post<RegisterTestUser>(APIURL + 'test/register', register, httpOptions).map(
      response => {
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
        });
  }

}
