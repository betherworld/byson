import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";


@Component({
  selector: 'app-donor-header',
  templateUrl: './donor-header.component.html',
  styleUrls: ['./donor-header.component.scss']
})
export class DonorHeaderComponent implements OnInit {
  isCollapsed = true;


  constructor(private router: Router) { }


  ngOnInit() {
  }


  navigate(donorEnvironment: string) {
    this.router.navigateByUrl(donorEnvironment);
  }
}
