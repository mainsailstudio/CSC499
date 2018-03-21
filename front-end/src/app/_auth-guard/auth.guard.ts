import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';

import { RedirectMessageService } from '../misc/redirect-message.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(private router: Router, private redirectMessage: RedirectMessageService) { }

    canActivate() {
        if (localStorage.getItem('currentUser')) {
            // logged in so return true
            return true;
        }

        // not logged in so redirect to login page
        this.redirectMessage.message = 'Please login first';
        this.router.navigate(['/']);
        return false;
    }
}
