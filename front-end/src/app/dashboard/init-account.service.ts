import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { APIURL, TestURL } from '../api/api.constants';
import { InitUser } from './dashboard.component';

@Injectable()
export class InitAccountService {

  private handleError: HandleError;

  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('InitAccountService');
  }

  initAccount (user: InitUser, token: string): Observable<InitUser> {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        'Authorization': 'Bearer ' + token
      })
    };

    return this.http.post<InitUser>(APIURL + 'register-continue', user, httpOptionsAuthorized)
      .pipe(
        catchError(this.handleError('initAccount', user))
      );
  }

  postPassword(userID: string, password: string, token: string) {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    interface UserPass {
      id: string;
      password: string;
    }

    // const hashArrayJSON = JSON.stringify(hashArray);
    // console.log('Hash array json is ' + hashArrayJSON);

    const userPass: UserPass = { id: userID, password: password } as UserPass;
    console.log('User auth before sending is ' + JSON.stringify(userPass));

    return this.http.post(APIURL + 'test/register-pass', userPass, httpOptionsAuthorized)
    .pipe(
      catchError(this.handleError('Register password', userPass))
    );
  }

  postKeys(userID: string, keys: string[], locks: string[], token: string) {
    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    interface UserKeys {
      id: string;
      keys:  string[];
      locks: string[];
    }

    const userKeys: UserKeys = { id: userID, keys: keys, locks: locks} as UserKeys;
    console.log('User auth before sending is ' + JSON.stringify(userKeys));

    return this.http.post(APIURL + 'test/register-keys', userKeys, httpOptionsAuthorized)
    .pipe(
      catchError(this.handleError('Register auth hash array', userKeys))
    );
  }


  postAuthArray(userID: string, locks: string[], hashArray: string[], token: string) {

    const httpOptionsAuthorized = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    interface UserAuth {
      id: string;
      locks:  string[];
      auths: string[];
    }

    // const hashArrayJSON = JSON.stringify(hashArray);
    // console.log('Hash array json is ' + hashArrayJSON);

    const userAuths: UserAuth = { id: userID, locks: locks, auths: hashArray } as UserAuth;
    console.log('User auth before sending is ' + JSON.stringify(userAuths));

    return this.http.post(APIURL + 'test/register-auth', userAuths, httpOptionsAuthorized)
    .pipe(
      catchError(this.handleError('Register auth hash array', userAuths))
    );
  }

}
