import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/timer';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/take';


import { AuthenticationService } from '../_auth-guard/authentication.service';

@Component({
    moduleId: module.id,
    templateUrl: 'logout.component.html'
})

export class LogoutComponent implements OnInit {
    countDown;
    counter = 4;

    constructor(private authenticationService: AuthenticationService, private router: Router) {}

    ngOnInit() {
        // Logout
        this.countDown = Observable.timer(0, 1000)
        .take(this.counter)
        .map(() => --this.counter);
        this.authenticationService.logout();
        Observable.timer(4200)
          .subscribe(i => {
            this.router.navigate(['/test']);
        });
    }
}
