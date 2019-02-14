import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { ProgressbarModule } from 'ngx-bootstrap/progressbar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BisonDashboardComponent } from './bison-dashboard/bison-dashboard.component';
import { DonorHeaderComponent } from './donor-header/donor-header.component';
import { RangerHeaderComponent } from './ranger-header/ranger-header.component';
import { RangerDashboardComponent } from './ranger-dashboard/ranger-dashboard.component';
import { RangerVotingComponent } from './ranger-voting/ranger-voting.component';
import { DonorImpactDashboardComponent } from './donor-impact-dashboard/donor-impact-dashboard.component';
import { DonorHomeComponent } from './donor-home/donor-home.component';
import { AgmCoreModule } from '@agm/core';
import { RangerProofComponent } from './ranger-proof/ranger-proof.component';
import { FormsModule } from "@angular/forms";

@NgModule({
  declarations: [
    AppComponent,
    BisonDashboardComponent,
    DonorHeaderComponent,
    RangerHeaderComponent,
    RangerDashboardComponent,
    RangerVotingComponent,
    DonorImpactDashboardComponent,
    DonorHomeComponent,
    RangerProofComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,

    CollapseModule.forRoot(),
    ProgressbarModule.forRoot(),
    AgmCoreModule.forRoot({
      apiKey: 'AIzaSyCaR1e6VL8iLDrJL0-NFN49WRCg3-_o-nY'
    }),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
