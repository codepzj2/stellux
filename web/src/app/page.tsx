
export default async function App() {
  await new Promise((resolve) => setTimeout(resolve, 4000));
  return (
    <div>
      <div className="text-red-500">hello</div>
    </div>
  );
}
