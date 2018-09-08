import { types } from "mobx-state-tree";

const Author = types.model({
  id: types.string,
  name: types.string,
  avatar: types.string
});

export default Author;
