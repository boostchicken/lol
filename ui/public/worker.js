
window.precacheAndRoute(self.__WB_MANIFEST);
registerRoute(
  ({url}) => url.pathname.startsWith('/liveconfig'),
  new StaleWhileRevalidate({cacheName: CACHE})
)
registerRoute(
  ({url}) => url.pathname.startsWith('/add'),
  new NetworkOnly()
);
registerRoute(
  ({url}) => url.pathname.startsWith('/lol'),
  new StaleWhileRevalidate({
    cacheName: CACHE
  })
);
registerRoute(
  ({url}) => url.pathname.startsWith('/del'),
  new NetworkOnly()
);
registerRoute(
  ({url}) => url.pathname.startsWith('/api'),
  new CacheFirst({
    cacheName: "lolapi"
  })
);
