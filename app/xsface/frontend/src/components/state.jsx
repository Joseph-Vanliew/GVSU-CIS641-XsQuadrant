import { atom } from "recoil";

export const videoStateAtom = atom({
  key: "videoState",
  default: "Stop Video",
});