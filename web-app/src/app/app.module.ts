import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CollapseModule } from 'ngx-bootstrap/collapse';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BisonDashboardComponent } from './bison-dashboard/bison-dashboard.component';
import { DonorHeaderComponent } from './donor-header/donor-header.component';
import { RangerHeaderComponent } from './ranger-header/ranger-header.component';
import { RangerDashboardComponent } from './ranger-dashboard/ranger-dashboard.component';
import { RangerVotingComponent } from './ranger-voting/ranger-voting.component';

@NgModule({
  declarations: [
    AppComponent,
    BisonDashboardComponent,
    DonorHeaderComponent,
    RangerHeaderComponent,
    RangerDashboardComponent,
    RangerVotingComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    CollapseModule.forRoot(),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
