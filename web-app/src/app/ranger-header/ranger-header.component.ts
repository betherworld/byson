import {Component, Input, OnInit} from '@angular/core';
import {Router} from "@angular/router";

@Component({
  selector: 'app-ranger-header',
  templateUrl: './ranger-header.component.html',
  styleUrls: ['./ranger-header.component.scss']
})
export class RangerHeaderComponent implements OnInit {
  @Input()
  siteNameBlack: string;

  @Input()
  siteNameGreen: string;

  @Input()
  showMenu: boolean = true;

  isCollapsed = true;


  constructor(private router: Router) { }

  ngOnInit() {
  }

  navigate(url: string) {
    this.router.navigateByUrl('/ranger/' + url);
    this.isCollapsed = true;
  }
}
