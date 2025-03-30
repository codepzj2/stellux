'use client';
 
import { ProgressProvider as NextProgressProvider } from '@bprogress/next/app';
 
export default function ProgressProvider({ children }: { children: React.ReactNode }) {
  return (
    <NextProgressProvider 
      height="2.5px"
      color="#539dfd"
      options={{ showSpinner: true }}
      shallowRouting
    >
      {children}
    </NextProgressProvider>
  );
};