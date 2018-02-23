import {Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {Observable} from 'rxjs/Observable';

const httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable()
export class TestService {

    constructor(private http: HttpClient) {}

    // Uses http.get() to load data from a single API endpoint
    getUsers() {
        // return this.http.get('http://13.92.156.114:8080');
        // return this.http.get('localhost:8080');
        return this.http.get('https://jsonplaceholder.typicode.com/posts/1');
    }
}
