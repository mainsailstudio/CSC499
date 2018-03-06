import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';

import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';
import { InitUser } from './dashboard.component';

@Injectable()
export class AccountInitService {

    userDataJSON = localStorage.getItem('currentUser');
    userData = JSON.parse(this.userDataJSON);
    token = this.userData['token'];
    authorization = 'Bearer ' + this.token;

    httpOptions = {
    headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.authorization,
        'Access-Control-Allow-Origin': '*',
        })
    };

    apiURL = 'http://localhost:8080/';
    private handleError: HandleError;

    constructor(
        private http: HttpClient,
        httpErrorHandler: HttpErrorHandler) {
        console.log('auth is ' + this.authorization);
        console.log('Token full is ' + this.token);
        this.handleError = httpErrorHandler.createHandleError('AccountInitService');
    }

    initAccount (user: InitUser): Observable<InitUser> {
        return this.http.post<InitUser>(this.apiURL + 'register-continue', user, this.httpOptions)
          .pipe(
            catchError(this.handleError('initAccount', user))
          );
    }
}
