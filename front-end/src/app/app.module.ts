import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';


import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { LoginStartComponent } from './login/login-start.component';
import { RegisterComponent } from './register/register.component';
import { RegisterStartComponent } from './register/register-start.component';
import { LoginService } from './login/login.service';
import { AppRoutingModule } from './app-routing.module';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { TestService } from './landing/test.service';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    LoginStartComponent,
    RegisterComponent,
    RegisterStartComponent,
    LandingComponent,
    NotFoundComponent
   ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [
    LoginService,
    TestService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
