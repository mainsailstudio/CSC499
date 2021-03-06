// package imports
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { HttpModule, BaseRequestOptions } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
// password strength bar
import { PasswordStrengthBar } from './misc/password-strength';

// app component imports
import { AppComponent } from './app.component';
import { LoginComponent, LoginSuccessComponent } from './login/login.component';
import { LoginNewComponent } from './login/login-new.component';
import { RegisterStartComponent, RegisterContinueComponent } from './register/register.component';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';
import {  DashboardComponent,
  DashboardMainComponent,
  DashboardInitComponent,
  DashboardSidebarComponent,
  DashboardNavComponent,
  DashboardPracticeComponent,
  DashboardHintsComponent,
  DashboardAboutComponent} from './dashboard/dashboard.component';
import { LogoutComponent } from './login/logout.component';

// module imports
import { AppRoutingModule } from './app-routing.module';

// service imports
import { AuthenticationService } from './_auth-guard/authentication.service';
import { AuthGuard } from './_auth-guard/auth.guard';
import { UserService } from './_auth-guard/user.service';
import { LoginUserService } from './login/login-user.service';
import { RegisterUserService } from './register/register-user.service';
import { HttpErrorHandler } from './http-error-handler.service';
import { MessageService } from './message.service';
import { InitAccountService } from './dashboard/init-account.service';
import { PermutateService } from './hash/perm.service';
import { HashSha256Service } from './hash/hash-sha256.service';
import { CombinePermsService } from './hash/combine.service';
import { RedirectMessageService } from './misc/redirect-message.service';


// test imports
import { LandingTestComponent } from './landing/landing-test.component';
import { RegisterTestComponent } from './register/register-test.component';
import { RegisterTestService } from './register/register-test.service';
import { LoginTestComponent } from './login/login-test.component';
import { LoginTestService } from './login/login-test.service';
import { DashboardTestComponent } from './dashboard/dashboard-test.component';
import { PracticeComponent } from './dashboard/practice/practice.component';
import { ActivityLogService } from './activity-log/activity-log.service';
import { PracticeService } from './dashboard/practice/practice.service';
import { UserConstantsService } from './dashboard/user-constants/user-constants.service';
import { AboutComponent } from './about/about.component';
import { HintsComponent } from './hints/hints.component';
import { UsabilityTestComponent, UsabilityTestInstructionsComponent } from './dashboard/usability-test/usability-test.component';
import { GetTokenService } from './misc/get-token.service';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    LoginSuccessComponent,
    LoginNewComponent,
    LogoutComponent,
    RegisterStartComponent,
    RegisterContinueComponent,
    LandingComponent,
    NotFoundComponent,
    DashboardComponent,
    DashboardMainComponent,
    DashboardInitComponent,
    DashboardSidebarComponent,
    DashboardNavComponent,
    RegisterTestComponent,
    LandingTestComponent,
    LoginTestComponent,
    DashboardTestComponent,
    PracticeComponent,
    DashboardPracticeComponent,
    AboutComponent,
    HintsComponent,
    DashboardHintsComponent,
    DashboardAboutComponent,
    PasswordStrengthBar,
    UsabilityTestComponent,
    UsabilityTestInstructionsComponent
  ],
  imports: [
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
    InitAccountService,
    LoginUserService,
    RegisterUserService,
    HttpErrorHandler,
    MessageService,
    BaseRequestOptions,
    RegisterTestService,
    LoginTestService,
    PermutateService,
    HashSha256Service,
    CombinePermsService,
    RedirectMessageService,
    ActivityLogService,
    PracticeService,
    UserConstantsService,
    GetTokenService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
