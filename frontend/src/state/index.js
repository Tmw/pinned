import { types } from "mobx-state-tree";

import { ChallangeStore, CreateChallangeStore } from "./ChallangeStore";
import { ViewStore, CreateViewStore } from "./ViewStore";

const Store = types.model({
  ViewStore,
  ChallangeStore
});

const DefaultState = {
  ViewStore: CreateViewStore(),
  ChallangeStore: CreateChallangeStore()
};

export default Store.create(DefaultState);
