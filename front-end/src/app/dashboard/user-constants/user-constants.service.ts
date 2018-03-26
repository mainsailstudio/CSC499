import { Injectable } from '@angular/core';

export interface UserConsts {
    ID: string;
    Email: string;
    Token: string;
    TestLevel: number;
    LoginState: number;
    Init: boolean;
}

@Injectable()
export class UserConstantsService {
    // this is where all of the user constants will be stored over a session
    // this shoud help cleanup the component code
    // and reduce the amount of times we need to rely on local storage
    // since local storage is unreliable

    public ID;
    public Email;
    public Token;
    public TestLevel;
    public LoginState;
    public Init;

    constructor() {
        const userDataJSON = localStorage.getItem('currentUser');
        const userData = JSON.parse(userDataJSON);

        if (userData == null) {
            return;
        }

        // everytime the service is called, check to make sure the variables aren't undefined
        // the localStorage is the final stage for making sure the information is correct
        if (this.ID === 'undefined' || this.ID === undefined || this.ID === '' || this.ID === null) {
            this.ID = userData['id'];
        }
        if (this.Email === 'undefined' || this.Email === undefined || this.Email === '' || this.Email === null) {
            this.Email = userData['email'];
        }
        if (this.Token === 'undefined' || this.Token === undefined || this.Token === '' || this.Token === null) {
            this.Token = userData['token'];
        }
        if (this.TestLevel === 'undefined' || this.TestLevel === undefined || this.TestLevel === '' || this.TestLevel === null) {
            this.TestLevel = userData['testLevel'];
        }
        if (this.LoginState === 'undefined' || this.LoginState === undefined || this.LoginState === '' || this.LoginState === null) {
            this.LoginState = userData['loginState'];
        }
        if (this.Init === 'undefined' || this.Init === undefined || this.Init === '' || this.Init === null) {
            this.Init = userData['init'];
        }
    }
}
