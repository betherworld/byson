import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";

@Component({
  selector: 'app-ranger-dashboard',
  templateUrl: './ranger-dashboard.component.html',
  styleUrls: ['./ranger-dashboard.component.scss']
})
export class RangerDashboardComponent implements OnInit {
  numberEther: number = 7.8065;
  numberRan: number = 25;

  constructor(private router: Router) { }

  ngOnInit() {
  }

  onProofClicked() {
    this.router.navigateByUrl('/ranger/proof');
  }
}
