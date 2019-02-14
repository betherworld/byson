import { Component, OnInit, ViewChild } from '@angular/core';


@Component({
  selector: 'app-ranger-proof',
  templateUrl: './ranger-proof.component.html',
  styleUrls: ['./ranger-proof.component.scss']
})
export class RangerProofComponent implements OnInit {

  @ViewChild('photo') file;
  photo;
  comment;


  constructor() { }

  ngOnInit() {
  }

  onPhotoAdded() {
    const files: FileList = this.file.nativeElement.files;
    if (files.length > 0) {
      this.photo = files[0];
    }
  }


  onSendClicked() {
    // TODO send
    console.log(this.photo, this.comment);
  }
}
