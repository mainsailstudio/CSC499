// package imports
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { HttpModule, BaseRequestOptions } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
// import { MDBBootstrapModule } from 'angular-bootstrap-md';

// app component imports
import { AppComponent } from './app.component';
import { LoginComponent, LoginSuccessComponent } from './login/login.component';
import { LoginNewComponent } from './login/login-new.component';
import { RegisterStartComponent, RegisterContinueComponent } from './register/register.component';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';

// module imports
import { AppRoutingModule } from './app-routing.module';

// service imports
// import { AuthenticateService } from './authenticate/authenticate.service';
import { AuthenticationService } from './_auth-guard/authentication.service';
import { AuthGuard } from './_auth-guard/auth.guard';
import { UserService } from './_auth-guard/user.service';
import { LoginUserService } from './login/login-user.service';
import { RegisterUserService } from './register/register-user.service';
import { HttpErrorHandler } from './http-error-handler.service';
import { MessageService } from './message.service';
import {  DashboardComponent,
          DashboardMainComponent,
          DashboardSidebarComponent,
          DashboardNavComponent } from './dashboard/dashboard.component';
import { ConfigurationComponent } from './dashboard/configuration/configuration.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    LoginSuccessComponent,
    LoginNewComponent,
    RegisterStartComponent,
    RegisterContinueComponent,
    LandingComponent,
    NotFoundComponent,
    DashboardComponent,
    DashboardMainComponent,
    DashboardSidebarComponent,
    DashboardNavComponent,
    ConfigurationComponent
   ],
  imports: [
    // MDBBootstrapModule.forRoot(),
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    HttpClientModule,
    HttpModule,
    FormsModule
  ],
  providers: [
    AuthenticationService,
    AuthGuard,
    UserService,
    LoginUserService,
    RegisterUserService,
    HttpErrorHandler,
    MessageService,
    BaseRequestOptions
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
