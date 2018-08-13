import { types } from "mobx-state-tree";

import { ChallengeStore, CreateChallengeStore } from "state/ChallengeStore";
import { ViewStore, CreateViewStore } from "state/ViewStore";

const Store = types.model({
  ViewStore,
  ChallengeStore
});

const DefaultState = {
  ViewStore: CreateViewStore(),
  ChallengeStore: CreateChallengeStore()
};

export default Store.create(DefaultState);
