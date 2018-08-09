import { types } from "mobx-state-tree";

import { ChallengeStore, CreateChallengeStore } from "./ChallengeStore";
import { ViewStore, CreateViewStore } from "./ViewStore";

const Store = types.model({
  ViewStore,
  ChallengeStore
});

const DefaultState = {
  ViewStore: CreateViewStore(),
  ChallengeStore: CreateChallengeStore()
};

export default Store.create(DefaultState);
