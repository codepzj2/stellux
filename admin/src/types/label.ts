export interface LabelVO {
  id: string;
  label_type: string;
  name: string;
}

export interface CreateLabelReq {
  label_type: string;
  name: string;
}

export interface EditLabelReq {
  id: string;
  label_type: string;
  name: string;
}

export interface LabelPageReq {
  page_no: number;
  page_size: number;
  label_type: string;
}
