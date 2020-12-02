package polyfill

var AbortController = &Polyfill{
	code: AbortController_Code,
	deps: []*Polyfill{  Fetch,  Object_defineProperty,  Object_getOwnPropertyDescriptor,  Object_getPrototypeOf,  Object_setPrototypeOf,  Event,  Map,  Symbol,  Symbol_toStringTag,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Array_from = &Polyfill{
	code: Array_from_Code,
	deps: []*Polyfill{  _ESAbstract_IsCallable,  _ESAbstract_CreateMethodProperty,  _ESAbstract_GetMethod,  Symbol_iterator,  _ESAbstract_IsConstructor,  _ESAbstract_Construct,  _ESAbstract_ArrayCreate,  _ESAbstract_GetIterator,  _ESAbstract_IteratorClose,  _ESAbstract_ToString,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_Call,  _ESAbstract_CreateDataPropertyOrThrow,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  Map,  Set,  },
	versions: LastVersionBeforeSupport{},
}

var Array_isArray = &Polyfill{
	code: Array_isArray_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsArray,  },
	versions: LastVersionBeforeSupport{},
}

var Array_of = &Polyfill{
	code: Array_of_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsConstructor,  _ESAbstract_Construct,  _ESAbstract_ArrayCreate,  _ESAbstract_ToString,  _ESAbstract_CreateDataPropertyOrThrow,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_ααiterator = &Polyfill{
	code: Array_prototype_ααiterator_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  Array_prototype_values,  Symbol_iterator,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_copyWithin = &Polyfill{
	code: Array_prototype_copyWithin_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_HasProperty,  _ESAbstract_ToInteger,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_entries = &Polyfill{
	code: Array_prototype_entries_Code,
	deps: []*Polyfill{  _ArrayIterator,  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_every = &Polyfill{
	code: Array_prototype_every_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_ToBoolean,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_fill = &Polyfill{
	code: Array_prototype_fill_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToInteger,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_filter = &Polyfill{
	code: Array_prototype_filter_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ArraySpeciesCreate,  _ESAbstract_HasProperty,  _ESAbstract_ToBoolean,  _ESAbstract_Call,  _ESAbstract_CreateDataPropertyOrThrow,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_find = &Polyfill{
	code: Array_prototype_find_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToBoolean,  _ESAbstract_Call,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_findIndex = &Polyfill{
	code: Array_prototype_findIndex_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToBoolean,  _ESAbstract_Call,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_flat = &Polyfill{
	code: Array_prototype_flat_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToInteger,  _ESAbstract_ArraySpeciesCreate,  _ESAbstract_FlattenIntoArray,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_flatMap = &Polyfill{
	code: Array_prototype_flatMap_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ArraySpeciesCreate,  _ESAbstract_FlattenIntoArray,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_forEach = &Polyfill{
	code: Array_prototype_forEach_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_HasProperty,  _ESAbstract_Call,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_includes = &Polyfill{
	code: Array_prototype_includes_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToInteger,  _ESAbstract_SameValueZero,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_indexOf = &Polyfill{
	code: Array_prototype_indexOf_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToInteger,  _ESAbstract_HasProperty,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_keys = &Polyfill{
	code: Array_prototype_keys_Code,
	deps: []*Polyfill{  _ArrayIterator,  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_lastIndexOf = &Polyfill{
	code: Array_prototype_lastIndexOf_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToInteger,  _ESAbstract_HasProperty,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_map = &Polyfill{
	code: Array_prototype_map_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ArraySpeciesCreate,  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_Call,  _ESAbstract_CreateDataPropertyOrThrow,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_reduce = &Polyfill{
	code: Array_prototype_reduce_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_reduceRight = &Polyfill{
	code: Array_prototype_reduceRight_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_some = &Polyfill{
	code: Array_prototype_some_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_ToBoolean,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_sort = &Polyfill{
	code: Array_prototype_sort_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsCallable,  Array_prototype_map,  },
	versions: LastVersionBeforeSupport{},
}

var Array_prototype_values = &Polyfill{
	code: Array_prototype_values_Code,
	deps: []*Polyfill{  _ArrayIterator,  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  },
	versions: LastVersionBeforeSupport{},
}

var AudioContext = &Polyfill{
	code: AudioContext_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Blob = &Polyfill{
	code: Blob_Code,
	deps: []*Polyfill{  Atob,  URL,  },
	versions: LastVersionBeforeSupport{},
}

var CSS_supports = &Polyfill{
	code: CSS_supports_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var CustomEvent = &Polyfill{
	code: CustomEvent_Code,
	deps: []*Polyfill{  Event,  },
	versions: LastVersionBeforeSupport{},
}

var DOMRect = &Polyfill{
	code: DOMRect_Code,
	deps: []*Polyfill{  Object_defineProperties,  },
	versions: LastVersionBeforeSupport{},
}

var DOMTokenList = &Polyfill{
	code: DOMTokenList_Code,
	deps: []*Polyfill{  _DOMTokenList,  },
	versions: LastVersionBeforeSupport{},
}

var DOMTokenList_prototype_ααiterator = &Polyfill{
	code: DOMTokenList_prototype_ααiterator_Code,
	deps: []*Polyfill{  _ArrayIterator,  Symbol_iterator,  DOMTokenList,  },
	versions: LastVersionBeforeSupport{},
}

var DOMTokenList_prototype_forEach = &Polyfill{
	code: DOMTokenList_prototype_forEach_Code,
	deps: []*Polyfill{  DOMTokenList,  Array_prototype_forEach,  Element_prototype_classList,  },
	versions: LastVersionBeforeSupport{},
}

var DOMTokenList_prototype_replace = &Polyfill{
	code: DOMTokenList_prototype_replace_Code,
	deps: []*Polyfill{  Element_prototype_classList,  },
	versions: LastVersionBeforeSupport{},
}

var Date_now = &Polyfill{
	code: Date_now_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Date_prototype_toISOString = &Polyfill{
	code: Date_prototype_toISOString_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var DocumentFragment = &Polyfill{
	code: DocumentFragment_Code,
	deps: []*Polyfill{  Object_create,  },
	versions: LastVersionBeforeSupport{},
}

var DocumentFragment_prototype_append = &Polyfill{
	code: DocumentFragment_prototype_append_Code,
	deps: []*Polyfill{  Document,  DocumentFragment,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var DocumentFragment_prototype_prepend = &Polyfill{
	code: DocumentFragment_prototype_prepend_Code,
	deps: []*Polyfill{  Document,  DocumentFragment,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element = &Polyfill{
	code: Element_Code,
	deps: []*Polyfill{  Document,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_after = &Polyfill{
	code: Element_prototype_after_Code,
	deps: []*Polyfill{  Document,  Element,  Array_prototype_indexOf,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_append = &Polyfill{
	code: Element_prototype_append_Code,
	deps: []*Polyfill{  Document,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_before = &Polyfill{
	code: Element_prototype_before_Code,
	deps: []*Polyfill{  Document,  Element,  Array_prototype_indexOf,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_classList = &Polyfill{
	code: Element_prototype_classList_Code,
	deps: []*Polyfill{  DOMTokenList,  Element,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_cloneNode = &Polyfill{
	code: Element_prototype_cloneNode_Code,
	deps: []*Polyfill{  Element,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_closest = &Polyfill{
	code: Element_prototype_closest_Code,
	deps: []*Polyfill{  Element_prototype_matches,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_dataset = &Polyfill{
	code: Element_prototype_dataset_Code,
	deps: []*Polyfill{  Object_defineProperty,  Object_getOwnPropertyDescriptor,  Document_querySelector,  Element,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_inert = &Polyfill{
	code: Element_prototype_inert_Code,
	deps: []*Polyfill{  Array_prototype_forEach,  Element,  Element_prototype_matches,  Function_prototype_bind,  Map,  MutationObserver,  Node_prototype_contains,  Object_defineProperty,  Set,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_matches = &Polyfill{
	code: Element_prototype_matches_Code,
	deps: []*Polyfill{  Element,  Document_querySelector,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_nextElementSibling = &Polyfill{
	code: Element_prototype_nextElementSibling_Code,
	deps: []*Polyfill{  Element,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_placeholder = &Polyfill{
	code: Element_prototype_placeholder_Code,
	deps: []*Polyfill{  Object_defineProperty,  Document_querySelector,  Element,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_prepend = &Polyfill{
	code: Element_prototype_prepend_Code,
	deps: []*Polyfill{  Document,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_previousElementSibling = &Polyfill{
	code: Element_prototype_previousElementSibling_Code,
	deps: []*Polyfill{  Element,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_remove = &Polyfill{
	code: Element_prototype_remove_Code,
	deps: []*Polyfill{  Document,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_replaceWith = &Polyfill{
	code: Element_prototype_replaceWith_Code,
	deps: []*Polyfill{  Document,  Element,  _mutation,  },
	versions: LastVersionBeforeSupport{},
}

var Element_prototype_toggleAttribute = &Polyfill{
	code: Element_prototype_toggleAttribute_Code,
	deps: []*Polyfill{  Element,  },
	versions: LastVersionBeforeSupport{},
}

var Event = &Polyfill{
	code: Event_Code,
	deps: []*Polyfill{  Window,  Document,  Element,  Object_defineProperty,  Array_prototype_indexOf,  Array_prototype_includes,  },
	versions: LastVersionBeforeSupport{},
}

var Event_focusin = &Polyfill{
	code: Event_focusin_Code,
	deps: []*Polyfill{  Event,  },
	versions: LastVersionBeforeSupport{},
}

var Event_hashchange = &Polyfill{
	code: Event_hashchange_Code,
	deps: []*Polyfill{  Event,  },
	versions: LastVersionBeforeSupport{},
}

var EventSource = &Polyfill{
	code: EventSource_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Function_prototype_bind = &Polyfill{
	code: Function_prototype_bind_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsCallable,  },
	versions: LastVersionBeforeSupport{},
}

var Function_prototype_name = &Polyfill{
	code: Function_prototype_name_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var HTMLCanvasElement_prototype_toBlob = &Polyfill{
	code: HTMLCanvasElement_prototype_toBlob_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _TypedArray,  Atob,  Blob,  },
	versions: LastVersionBeforeSupport{},
}

var HTMLDocument = &Polyfill{
	code: HTMLDocument_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var HTMLPictureElement = &Polyfill{
	code: HTMLPictureElement_Code,
	deps: []*Polyfill{  __html5_elements,  },
	versions: LastVersionBeforeSupport{},
}

var HTMLTemplateElement = &Polyfill{
	code: HTMLTemplateElement_Code,
	deps: []*Polyfill{  DocumentFragment,  Document_querySelector,  Object_defineProperties,  Object_getOwnPropertyDescriptor,  Object_defineProperty,  Event,  },
	versions: LastVersionBeforeSupport{},
}

var IntersectionObserver = &Polyfill{
	code: IntersectionObserver_Code,
	deps: []*Polyfill{  GetComputedStyle,  Array_isArray,  Array_prototype_filter,  Array_prototype_forEach,  Array_prototype_indexOf,  Array_prototype_map,  Array_prototype_some,  Event,  Function_prototype_bind,  Performance_now,  },
	versions: LastVersionBeforeSupport{},
}

var IntersectionObserverEntry = &Polyfill{
	code: IntersectionObserverEntry_Code,
	deps: []*Polyfill{  IntersectionObserver,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_DateTimeFormat = &Polyfill{
	code: Intl_DateTimeFormat_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Date_now,  Intl_getCanonicalLocales,  Intl_Locale,  Intl_NumberFormat,  Object_assign,  Object_create,  Object_keys,  Object_setPrototypeOf,  Set,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_DisplayNames = &Polyfill{
	code: Intl_DisplayNames_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Intl_getCanonicalLocales,  Intl_Locale,  Object_assign,  Object_create,  Object_defineProperty,  Object_keys,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_ListFormat = &Polyfill{
	code: Intl_ListFormat_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_forEach,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Intl_getCanonicalLocales,  Intl_Locale,  Object_assign,  Object_create,  Object_defineProperty,  Object_keys,  Object_setPrototypeOf,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_Locale = &Polyfill{
	code: Intl_Locale_Code,
	deps: []*Polyfill{  Array_prototype_filter,  Array_prototype_indexOf,  Array_prototype_reduce,  Array_prototype_sort,  Intl_getCanonicalLocales,  Object_assign,  Object_defineProperty,  Object_keys,  Set,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_NumberFormat = &Polyfill{
	code: Intl_NumberFormat_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_forEach,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Intl_getCanonicalLocales,  Intl_Locale,  Intl_PluralRules,  Object_assign,  Object_create,  Object_defineProperty,  Object_keys,  Object_setPrototypeOf,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_PluralRules = &Polyfill{
	code: Intl_PluralRules_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_forEach,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Intl_getCanonicalLocales,  Intl_Locale,  Object_assign,  Object_create,  Object_defineProperty,  Object_keys,  Object_setPrototypeOf,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_RelativeTimeFormat = &Polyfill{
	code: Intl_RelativeTimeFormat_Code,
	deps: []*Polyfill{  Array_isArray,  Array_prototype_filter,  Array_prototype_forEach,  Array_prototype_indexOf,  Array_prototype_lastIndexOf,  Array_prototype_map,  Array_prototype_reduce,  Intl_getCanonicalLocales,  Intl_Locale,  Intl_NumberFormat,  Intl_PluralRules,  Object_assign,  Object_create,  Object_defineProperty,  Object_keys,  Object_setPrototypeOf,  WeakMap,  },
	versions: LastVersionBeforeSupport{},
}

var Intl_getCanonicalLocales = &Polyfill{
	code: Intl_getCanonicalLocales_Code,
	deps: []*Polyfill{  Array_prototype_filter,  Array_prototype_indexOf,  Array_prototype_reduce,  Array_prototype_sort,  Object_assign,  Object_defineProperty,  Object_keys,  },
	versions: LastVersionBeforeSupport{},
}

var JSON = &Polyfill{
	code: JSON_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Map = &Polyfill{
	code: Map_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_CreateIterResultObject,  _ESAbstract_GetMethod,  _ESAbstract_GetIterator,  _ESAbstract_IsCallable,  _ESAbstract_IteratorClose,  _ESAbstract_IteratorComplete,  _ESAbstract_IteratorNext,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_OrdinaryCreateFromConstructor,  _ESAbstract_SameValueZero,  _ESAbstract_Type,  Array_isArray,  Symbol,  Symbol_iterator,  Symbol_species,  Object_create,  Object_defineProperty,  Object_isExtensible,  },
	versions: LastVersionBeforeSupport{},
}

var Math_acosh = &Polyfill{
	code: Math_acosh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_asinh = &Polyfill{
	code: Math_asinh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_atanh = &Polyfill{
	code: Math_atanh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_cbrt = &Polyfill{
	code: Math_cbrt_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_clz32 = &Polyfill{
	code: Math_clz32_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToUint32,  },
	versions: LastVersionBeforeSupport{},
}

var Math_cosh = &Polyfill{
	code: Math_cosh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_expm1 = &Polyfill{
	code: Math_expm1_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_fround = &Polyfill{
	code: Math_fround_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _TypedArray,  },
	versions: LastVersionBeforeSupport{},
}

var Math_hypot = &Polyfill{
	code: Math_hypot_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_imul = &Polyfill{
	code: Math_imul_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToUint32,  },
	versions: LastVersionBeforeSupport{},
}

var Math_log10 = &Polyfill{
	code: Math_log10_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_log1p = &Polyfill{
	code: Math_log1p_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_log2 = &Polyfill{
	code: Math_log2_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_sign = &Polyfill{
	code: Math_sign_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_sinh = &Polyfill{
	code: Math_sinh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_tanh = &Polyfill{
	code: Math_tanh_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Math_trunc = &Polyfill{
	code: Math_trunc_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var MutationObserver = &Polyfill{
	code: MutationObserver_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Node_prototype_contains = &Polyfill{
	code: Node_prototype_contains_Code,
	deps: []*Polyfill{  Element,  },
	versions: LastVersionBeforeSupport{},
}

var NodeList_prototype_ααiterator = &Polyfill{
	code: NodeList_prototype_ααiterator_Code,
	deps: []*Polyfill{  _ArrayIterator,  Symbol_iterator,  },
	versions: LastVersionBeforeSupport{},
}

var NodeList_prototype_forEach = &Polyfill{
	code: NodeList_prototype_forEach_Code,
	deps: []*Polyfill{  Array_prototype_forEach,  },
	versions: LastVersionBeforeSupport{},
}

var Number_Epsilon = &Polyfill{
	code: Number_Epsilon_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Number_MAX_SAFE_INTEGER = &Polyfill{
	code: Number_MAX_SAFE_INTEGER_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Number_MIN_SAFE_INTEGER = &Polyfill{
	code: Number_MIN_SAFE_INTEGER_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Number_isFinite = &Polyfill{
	code: Number_isFinite_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Number_isInteger = &Polyfill{
	code: Number_isInteger_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToInteger,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Number_isNaN = &Polyfill{
	code: Number_isNaN_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Number_isSafeInteger = &Polyfill{
	code: Number_isSafeInteger_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToInteger,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Number_parseFloat = &Polyfill{
	code: Number_parseFloat_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  String_prototype_trim,  },
	versions: LastVersionBeforeSupport{},
}

var Number_parseInt = &Polyfill{
	code: Number_parseInt_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  String_prototype_trim,  },
	versions: LastVersionBeforeSupport{},
}

var Object_assign = &Polyfill{
	code: Object_assign_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Get,  _ESAbstract_ToObject,  Object_getOwnPropertyDescriptor,  Object_keys,  },
	versions: LastVersionBeforeSupport{},
}

var Object_create = &Polyfill{
	code: Object_create_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  Object_defineProperties,  },
	versions: LastVersionBeforeSupport{},
}

var Object_defineProperties = &Polyfill{
	code: Object_defineProperties_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Get,  _ESAbstract_ToObject,  _ESAbstract_Type,  Object_keys,  Object_getOwnPropertyDescriptor,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Object_defineProperty = &Polyfill{
	code: Object_defineProperty_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Object_entries = &Polyfill{
	code: Object_entries_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_EnumerableOwnProperties,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_freeze = &Polyfill{
	code: Object_freeze_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Object_fromEntries = &Polyfill{
	code: Object_fromEntries_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToPropertyKey,  _ESAbstract_CreateDataPropertyOrThrow,  _ESAbstract_AddEntriesFromIterable,  Array_prototype_ααiterator,  },
	versions: LastVersionBeforeSupport{},
}

var Object_getOwnPropertyDescriptor = &Polyfill{
	code: Object_getOwnPropertyDescriptor_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_HasOwnProperty,  _ESAbstract_ToObject,  _ESAbstract_ToPropertyKey,  _ESAbstract_Type,  Function_prototype_bind,  },
	versions: LastVersionBeforeSupport{},
}

var Object_getOwnPropertyDescriptors = &Polyfill{
	code: Object_getOwnPropertyDescriptors_Code,
	deps: []*Polyfill{  Object_getOwnPropertyDescriptor,  Object_defineProperty,  Reflect_ownKeys,  _ESAbstract_ToObject,  _ESAbstract_CreateMethodProperty,  _ESAbstract_CreateDataProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Object_getOwnPropertyNames = &Polyfill{
	code: Object_getOwnPropertyNames_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  Object_keys,  Array_prototype_includes,  Array_prototype_indexOf,  },
	versions: LastVersionBeforeSupport{},
}

var Object_getPrototypeOf = &Polyfill{
	code: Object_getPrototypeOf_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Object_is = &Polyfill{
	code: Object_is_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_SameValue,  },
	versions: LastVersionBeforeSupport{},
}

var Object_isExtensible = &Polyfill{
	code: Object_isExtensible_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_isFrozen = &Polyfill{
	code: Object_isFrozen_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_isSealed = &Polyfill{
	code: Object_isSealed_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_keys = &Polyfill{
	code: Object_keys_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Object_preventExtensions = &Polyfill{
	code: Object_preventExtensions_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_seal = &Polyfill{
	code: Object_seal_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Object_setPrototypeOf = &Polyfill{
	code: Object_setPrototypeOf_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  Array_prototype_forEach,  Object_create,  Object_defineProperty,  Object_getPrototypeOf,  Object_getOwnPropertyNames,  Object_getOwnPropertyDescriptor,  },
	versions: LastVersionBeforeSupport{},
}

var Object_values = &Polyfill{
	code: Object_values_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  Array_prototype_map,  Object_defineProperty,  Object_keys,  },
	versions: LastVersionBeforeSupport{},
}

var Promise = &Polyfill{
	code: Promise_Code,
	deps: []*Polyfill{  Symbol_toStringTag,  },
	versions: LastVersionBeforeSupport{},
}

var Promise_prototype_finally = &Polyfill{
	code: Promise_prototype_finally_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsCallable,  _ESAbstract_SpeciesConstructor,  Promise,  Function_prototype_bind,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect = &Polyfill{
	code: Reflect_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_apply = &Polyfill{
	code: Reflect_apply_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsCallable,  _ESAbstract_CreateListFromArrayLike,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_construct = &Polyfill{
	code: Reflect_construct_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_IsConstructor,  _ESAbstract_CreateListFromArrayLike,  _ESAbstract_Construct,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_defineProperty = &Polyfill{
	code: Reflect_defineProperty_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  _ESAbstract_ToPropertyDescriptor,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_deleteProperty = &Polyfill{
	code: Reflect_deleteProperty_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  _ESAbstract_HasOwnProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_get = &Polyfill{
	code: Reflect_get_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_Call,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  Object_getOwnPropertyDescriptor,  Object_getPrototypeOf,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_getOwnPropertyDescriptor = &Polyfill{
	code: Reflect_getOwnPropertyDescriptor_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  Object_getOwnPropertyDescriptor,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_getPrototypeOf = &Polyfill{
	code: Reflect_getPrototypeOf_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  Object_getPrototypeOf,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_has = &Polyfill{
	code: Reflect_has_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  _ESAbstract_HasProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_isExtensible = &Polyfill{
	code: Reflect_isExtensible_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  Object_isExtensible,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_ownKeys = &Polyfill{
	code: Reflect_ownKeys_Code,
	deps: []*Polyfill{  Reflect,  Symbol,  Object_getOwnPropertyNames,  _ESAbstract_Type,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_preventExtensions = &Polyfill{
	code: Reflect_preventExtensions_Code,
	deps: []*Polyfill{  Reflect,  Object_preventExtensions,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_set = &Polyfill{
	code: Reflect_set_Code,
	deps: []*Polyfill{  Reflect,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  _ESAbstract_ToPropertyKey,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var Reflect_setPrototypeOf = &Polyfill{
	code: Reflect_setPrototypeOf_Code,
	deps: []*Polyfill{  Reflect,  Reflect_getPrototypeOf,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Type,  Object_setPrototypeOf,  },
	versions: LastVersionBeforeSupport{},
}

var RegExp_prototype_flags = &Polyfill{
	code: RegExp_prototype_flags_Code,
	deps: []*Polyfill{  _ESAbstract_Get,  _ESAbstract_ToBoolean,  _ESAbstract_Type,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var ResizeObserver = &Polyfill{
	code: ResizeObserver_Code,
	deps: []*Polyfill{  WeakMap,  MutationObserver,  RequestAnimationFrame,  DevicePixelRatio,  },
	versions: LastVersionBeforeSupport{},
}

var Set = &Polyfill{
	code: Set_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_CreateIterResultObject,  _ESAbstract_GetMethod,  _ESAbstract_GetIterator,  _ESAbstract_IsCallable,  _ESAbstract_IteratorClose,  _ESAbstract_IteratorComplete,  _ESAbstract_IteratorNext,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_OrdinaryCreateFromConstructor,  _ESAbstract_SameValueZero,  Array_isArray,  Symbol,  Symbol_iterator,  Symbol_species,  Object_create,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_fromCodePoint = &Polyfill{
	code: String_fromCodePoint_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToNumber,  _ESAbstract_SameValue,  _ESAbstract_ToInteger,  _ESAbstract_IsArray,  _ESAbstract_UTF16Encoding,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_ααiterator = &Polyfill{
	code: String_prototype_ααiterator_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _StringIterator,  Symbol_iterator,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_anchor = &Polyfill{
	code: String_prototype_anchor_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_big = &Polyfill{
	code: String_prototype_big_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_blink = &Polyfill{
	code: String_prototype_blink_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_bold = &Polyfill{
	code: String_prototype_bold_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_codePointAt = &Polyfill{
	code: String_prototype_codePointAt_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_ToInteger,  _ESAbstract_UTF16Decode,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_endsWith = &Polyfill{
	code: String_prototype_endsWith_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_IsRegExp,  _ESAbstract_ToInteger,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_fixed = &Polyfill{
	code: String_prototype_fixed_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_fontcolor = &Polyfill{
	code: String_prototype_fontcolor_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_fontsize = &Polyfill{
	code: String_prototype_fontsize_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_includes = &Polyfill{
	code: String_prototype_includes_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_IsRegExp,  _ESAbstract_ToInteger,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_italics = &Polyfill{
	code: String_prototype_italics_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_link = &Polyfill{
	code: String_prototype_link_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_normalize = &Polyfill{
	code: String_prototype_normalize_Code,
	deps: []*Polyfill{  Array_prototype_reduceRight,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_padEnd = &Polyfill{
	code: String_prototype_padEnd_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_ToLength,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_padStart = &Polyfill{
	code: String_prototype_padStart_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_ToLength,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_repeat = &Polyfill{
	code: String_prototype_repeat_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToInteger,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_replaceAll = &Polyfill{
	code: String_prototype_replaceAll_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_CreateMethodProperty,  _ESAbstract_Get,  _ESAbstract_GetMethod,  _ESAbstract_GetSubstitution,  _ESAbstract_IsCallable,  _ESAbstract_IsRegExp,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_StringIndexOf,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_small = &Polyfill{
	code: String_prototype_small_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_startsWith = &Polyfill{
	code: String_prototype_startsWith_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  _ESAbstract_IsRegExp,  _ESAbstract_ToInteger,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_strike = &Polyfill{
	code: String_prototype_strike_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_sub = &Polyfill{
	code: String_prototype_sub_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_sup = &Polyfill{
	code: String_prototype_sup_Code,
	deps: []*Polyfill{  _ESAbstract_CreateHTML,  _ESAbstract_CreateMethodProperty,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_trim = &Polyfill{
	code: String_prototype_trim_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_TrimString,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_trimEnd = &Polyfill{
	code: String_prototype_trimEnd_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_TrimString,  },
	versions: LastVersionBeforeSupport{},
}

var String_prototype_trimStart = &Polyfill{
	code: String_prototype_trimStart_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_TrimString,  },
	versions: LastVersionBeforeSupport{},
}

var String_raw = &Polyfill{
	code: String_raw_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_ToObject,  _ESAbstract_Get,  _ESAbstract_ToLength,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol = &Polyfill{
	code: Symbol_Code,
	deps: []*Polyfill{  Array_prototype_forEach,  Array_prototype_filter,  Array_prototype_map,  Object_create,  Object_defineProperty,  Object_getOwnPropertyNames,  Object_getOwnPropertyDescriptor,  Object_freeze,  Object_keys,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_asyncIterator = &Polyfill{
	code: Symbol_asyncIterator_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_hasInstance = &Polyfill{
	code: Symbol_hasInstance_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_isConcatSpreadable = &Polyfill{
	code: Symbol_isConcatSpreadable_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_iterator = &Polyfill{
	code: Symbol_iterator_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_match = &Polyfill{
	code: Symbol_match_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_prototype_description = &Polyfill{
	code: Symbol_prototype_description_Code,
	deps: []*Polyfill{  Symbol,  Object_defineProperty,  String_prototype_trim,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_replace = &Polyfill{
	code: Symbol_replace_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_search = &Polyfill{
	code: Symbol_search_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_species = &Polyfill{
	code: Symbol_species_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_split = &Polyfill{
	code: Symbol_split_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_toPrimitive = &Polyfill{
	code: Symbol_toPrimitive_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_toStringTag = &Polyfill{
	code: Symbol_toStringTag_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var Symbol_unscopables = &Polyfill{
	code: Symbol_unscopables_Code,
	deps: []*Polyfill{  Object_defineProperty,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var TextEncoder = &Polyfill{
	code: TextEncoder_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var URL = &Polyfill{
	code: URL_Code,
	deps: []*Polyfill{  Object_defineProperties,  Array_prototype_forEach,  Array_isArray,  },
	versions: LastVersionBeforeSupport{},
}

var UserTiming = &Polyfill{
	code: UserTiming_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var WeakMap = &Polyfill{
	code: WeakMap_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_OrdinaryCreateFromConstructor,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_GetIterator,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_Call,  _ESAbstract_IteratorClose,  _ESAbstract_IsArray,  Object_defineProperty,  _ESAbstract_Type,  _ESAbstract_SameValue,  Symbol,  Array_isArray,  },
	versions: LastVersionBeforeSupport{},
}

var WeakSet = &Polyfill{
	code: WeakSet_Code,
	deps: []*Polyfill{  _ESAbstract_CreateMethodProperty,  _ESAbstract_OrdinaryCreateFromConstructor,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_GetIterator,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_Call,  _ESAbstract_IteratorClose,  _ESAbstract_IsArray,  Object_defineProperty,  _ESAbstract_Type,  _ESAbstract_SameValueZero,  Symbol,  },
	versions: LastVersionBeforeSupport{},
}

var WebAnimations = &Polyfill{
	code: WebAnimations_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Window = &Polyfill{
	code: Window_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var XMLHttpRequest = &Polyfill{
	code: XMLHttpRequest_Code,
	deps: []*Polyfill{  Event,  },
	versions: LastVersionBeforeSupport{},
}

var _ArrayIterator = &Polyfill{
	code: _ArrayIterator_Code,
	deps: []*Polyfill{  _Iterator,  Object_create,  Object_defineProperty,  Object_setPrototypeOf,  String_prototype_includes,  },
	versions: LastVersionBeforeSupport{},
}

var _DOMTokenList = &Polyfill{
	code: _DOMTokenList_Code,
	deps: []*Polyfill{  Object_defineProperty,  Array_prototype_forEach,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_AddEntriesFromIterable = &Polyfill{
	code: _ESAbstract_AddEntriesFromIterable_Code,
	deps: []*Polyfill{  _ESAbstract_IsCallable,  _ESAbstract_GetIterator,  _ESAbstract_IteratorStep,  _ESAbstract_IteratorValue,  _ESAbstract_IteratorClose,  _ESAbstract_Get,  _ESAbstract_Call,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ArrayCreate = &Polyfill{
	code: _ESAbstract_ArrayCreate_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ArraySpeciesCreate = &Polyfill{
	code: _ESAbstract_ArraySpeciesCreate_Code,
	deps: []*Polyfill{  _ESAbstract_IsArray,  _ESAbstract_ArrayCreate,  _ESAbstract_Get,  _ESAbstract_Type,  _ESAbstract_IsConstructor,  _ESAbstract_Construct,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_Call = &Polyfill{
	code: _ESAbstract_Call_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CanonicalNumericIndexString = &Polyfill{
	code: _ESAbstract_CanonicalNumericIndexString_Code,
	deps: []*Polyfill{  _ESAbstract_SameValue,  _ESAbstract_ToNumber,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_Construct = &Polyfill{
	code: _ESAbstract_Construct_Code,
	deps: []*Polyfill{  _ESAbstract_IsConstructor,  _ESAbstract_OrdinaryCreateFromConstructor,  Function_prototype_bind,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateDataProperty = &Polyfill{
	code: _ESAbstract_CreateDataProperty_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateDataPropertyOrThrow = &Polyfill{
	code: _ESAbstract_CreateDataPropertyOrThrow_Code,
	deps: []*Polyfill{  _ESAbstract_CreateDataProperty,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateHTML = &Polyfill{
	code: _ESAbstract_CreateHTML_Code,
	deps: []*Polyfill{  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateIterResultObject = &Polyfill{
	code: _ESAbstract_CreateIterResultObject_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_CreateDataProperty,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateListFromArrayLike = &Polyfill{
	code: _ESAbstract_CreateListFromArrayLike_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_ToLength,  _ESAbstract_Get,  _ESAbstract_ToString,  Array_prototype_includes,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_CreateMethodProperty = &Polyfill{
	code: _ESAbstract_CreateMethodProperty_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_EnumerableOwnProperties = &Polyfill{
	code: _ESAbstract_EnumerableOwnProperties_Code,
	deps: []*Polyfill{  _ESAbstract_Get,  _ESAbstract_Type,  Object_getOwnPropertyDescriptor,  Object_keys,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_FlattenIntoArray = &Polyfill{
	code: _ESAbstract_FlattenIntoArray_Code,
	deps: []*Polyfill{  _ESAbstract_ToString,  _ESAbstract_HasProperty,  _ESAbstract_Get,  _ESAbstract_Call,  _ESAbstract_IsArray,  _ESAbstract_ToLength,  _ESAbstract_CreateDataPropertyOrThrow,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_Get = &Polyfill{
	code: _ESAbstract_Get_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_GetIterator = &Polyfill{
	code: _ESAbstract_GetIterator_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_GetMethod,  _ESAbstract_GetV,  _ESAbstract_Type,  Object_create,  Symbol_iterator,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_GetMethod = &Polyfill{
	code: _ESAbstract_GetMethod_Code,
	deps: []*Polyfill{  _ESAbstract_GetV,  _ESAbstract_IsCallable,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_GetPrototypeFromConstructor = &Polyfill{
	code: _ESAbstract_GetPrototypeFromConstructor_Code,
	deps: []*Polyfill{  _ESAbstract_Get,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_GetSubstitution = &Polyfill{
	code: _ESAbstract_GetSubstitution_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_GetV = &Polyfill{
	code: _ESAbstract_GetV_Code,
	deps: []*Polyfill{  _ESAbstract_ToObject,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_HasOwnProperty = &Polyfill{
	code: _ESAbstract_HasOwnProperty_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_HasProperty = &Polyfill{
	code: _ESAbstract_HasProperty_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_Invoke = &Polyfill{
	code: _ESAbstract_Invoke_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_GetV,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsArray = &Polyfill{
	code: _ESAbstract_IsArray_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsCallable = &Polyfill{
	code: _ESAbstract_IsCallable_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsConstructor = &Polyfill{
	code: _ESAbstract_IsConstructor_Code,
	deps: []*Polyfill{  _ESAbstract_GetMethod,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsInteger = &Polyfill{
	code: _ESAbstract_IsInteger_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsPropertyKey = &Polyfill{
	code: _ESAbstract_IsPropertyKey_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IsRegExp = &Polyfill{
	code: _ESAbstract_IsRegExp_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_ToBoolean,  _ESAbstract_Get,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IteratorClose = &Polyfill{
	code: _ESAbstract_IteratorClose_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_GetMethod,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IteratorComplete = &Polyfill{
	code: _ESAbstract_IteratorComplete_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_ToBoolean,  _ESAbstract_Get,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IteratorNext = &Polyfill{
	code: _ESAbstract_IteratorNext_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_Call,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IteratorStep = &Polyfill{
	code: _ESAbstract_IteratorStep_Code,
	deps: []*Polyfill{  _ESAbstract_IteratorNext,  _ESAbstract_IteratorComplete,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_IteratorValue = &Polyfill{
	code: _ESAbstract_IteratorValue_Code,
	deps: []*Polyfill{  _ESAbstract_Get,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_OrdinaryCreateFromConstructor = &Polyfill{
	code: _ESAbstract_OrdinaryCreateFromConstructor_Code,
	deps: []*Polyfill{  _ESAbstract_GetPrototypeFromConstructor,  Object_create,  Object_defineProperty,  Object_getPrototypeOf,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_OrdinaryToPrimitive = &Polyfill{
	code: _ESAbstract_OrdinaryToPrimitive_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_Get,  _ESAbstract_IsCallable,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_RequireObjectCoercible = &Polyfill{
	code: _ESAbstract_RequireObjectCoercible_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_SameValue = &Polyfill{
	code: _ESAbstract_SameValue_Code,
	deps: []*Polyfill{  _ESAbstract_SameValueNonNumber,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_SameValueNonNumber = &Polyfill{
	code: _ESAbstract_SameValueNonNumber_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_SameValueZero = &Polyfill{
	code: _ESAbstract_SameValueZero_Code,
	deps: []*Polyfill{  _ESAbstract_SameValueNonNumber,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_SpeciesConstructor = &Polyfill{
	code: _ESAbstract_SpeciesConstructor_Code,
	deps: []*Polyfill{  _ESAbstract_Get,  _ESAbstract_IsConstructor,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_StringIndexOf = &Polyfill{
	code: _ESAbstract_StringIndexOf_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToBoolean = &Polyfill{
	code: _ESAbstract_ToBoolean_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToIndex = &Polyfill{
	code: _ESAbstract_ToIndex_Code,
	deps: []*Polyfill{  _ESAbstract_ToInteger,  _ESAbstract_ToLength,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToInt16 = &Polyfill{
	code: _ESAbstract_ToInt16_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToInt32 = &Polyfill{
	code: _ESAbstract_ToInt32_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToInt8 = &Polyfill{
	code: _ESAbstract_ToInt8_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToInteger = &Polyfill{
	code: _ESAbstract_ToInteger_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToLength = &Polyfill{
	code: _ESAbstract_ToLength_Code,
	deps: []*Polyfill{  _ESAbstract_ToInteger,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToNumber = &Polyfill{
	code: _ESAbstract_ToNumber_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToObject = &Polyfill{
	code: _ESAbstract_ToObject_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToPrimitive = &Polyfill{
	code: _ESAbstract_ToPrimitive_Code,
	deps: []*Polyfill{  _ESAbstract_Call,  _ESAbstract_GetMethod,  _ESAbstract_OrdinaryToPrimitive,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToPropertyDescriptor = &Polyfill{
	code: _ESAbstract_ToPropertyDescriptor_Code,
	deps: []*Polyfill{  _ESAbstract_Type,  _ESAbstract_HasProperty,  _ESAbstract_ToBoolean,  _ESAbstract_Get,  _ESAbstract_IsCallable,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToPropertyKey = &Polyfill{
	code: _ESAbstract_ToPropertyKey_Code,
	deps: []*Polyfill{  _ESAbstract_ToPrimitive,  _ESAbstract_ToString,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToString = &Polyfill{
	code: _ESAbstract_ToString_Code,
	deps: []*Polyfill{  _ESAbstract_ToPrimitive,  _ESAbstract_Type,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToUint16 = &Polyfill{
	code: _ESAbstract_ToUint16_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToUint32 = &Polyfill{
	code: _ESAbstract_ToUint32_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToUint8 = &Polyfill{
	code: _ESAbstract_ToUint8_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_ToUint8Clamp = &Polyfill{
	code: _ESAbstract_ToUint8Clamp_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_TrimString = &Polyfill{
	code: _ESAbstract_TrimString_Code,
	deps: []*Polyfill{  _ESAbstract_RequireObjectCoercible,  _ESAbstract_ToString,  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_Type = &Polyfill{
	code: _ESAbstract_Type_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_UTF16Decode = &Polyfill{
	code: _ESAbstract_UTF16Decode_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _ESAbstract_UTF16Encoding = &Polyfill{
	code: _ESAbstract_UTF16Encoding_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _Iterator = &Polyfill{
	code: _Iterator_Code,
	deps: []*Polyfill{  Function_prototype_bind,  Object_assign,  Object_defineProperties,  Object_defineProperty,  Symbol,  Symbol_iterator,  Symbol_toStringTag,  },
	versions: LastVersionBeforeSupport{},
}

var _StringIterator = &Polyfill{
	code: _StringIterator_Code,
	deps: []*Polyfill{  _Iterator,  Object_create,  Object_defineProperty,  Object_setPrototypeOf,  },
	versions: LastVersionBeforeSupport{},
}

var _TypedArray = &Polyfill{
	code: _TypedArray_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var _mutation = &Polyfill{
	code: _mutation_Code,
	deps: []*Polyfill{  Document,  Element,  },
	versions: LastVersionBeforeSupport{},
}

var Atob = &Polyfill{
	code: Atob_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Console = &Polyfill{
	code: Console_Code,
	deps: []*Polyfill{  Window,  },
	versions: LastVersionBeforeSupport{},
}

var Console_assert = &Polyfill{
	code: Console_assert_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_clear = &Polyfill{
	code: Console_clear_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_count = &Polyfill{
	code: Console_count_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_debug = &Polyfill{
	code: Console_debug_Code,
	deps: []*Polyfill{  Window,  Console,  Console_log,  },
	versions: LastVersionBeforeSupport{},
}

var Console_dir = &Polyfill{
	code: Console_dir_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_dirxml = &Polyfill{
	code: Console_dirxml_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_error = &Polyfill{
	code: Console_error_Code,
	deps: []*Polyfill{  Window,  Console,  Console_log,  },
	versions: LastVersionBeforeSupport{},
}

var Console_exception = &Polyfill{
	code: Console_exception_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_group = &Polyfill{
	code: Console_group_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_groupCollapsed = &Polyfill{
	code: Console_groupCollapsed_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_groupEnd = &Polyfill{
	code: Console_groupEnd_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_info = &Polyfill{
	code: Console_info_Code,
	deps: []*Polyfill{  Window,  Console,  Console_log,  },
	versions: LastVersionBeforeSupport{},
}

var Console_log = &Polyfill{
	code: Console_log_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_markTimeline = &Polyfill{
	code: Console_markTimeline_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_profile = &Polyfill{
	code: Console_profile_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_profileEnd = &Polyfill{
	code: Console_profileEnd_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_profiles = &Polyfill{
	code: Console_profiles_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_table = &Polyfill{
	code: Console_table_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_time = &Polyfill{
	code: Console_time_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_timeEnd = &Polyfill{
	code: Console_timeEnd_Code,
	deps: []*Polyfill{  Window,  Console,  Console_time,  },
	versions: LastVersionBeforeSupport{},
}

var Console_timeStamp = &Polyfill{
	code: Console_timeStamp_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_timeline = &Polyfill{
	code: Console_timeline_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_timelineEnd = &Polyfill{
	code: Console_timelineEnd_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_trace = &Polyfill{
	code: Console_trace_Code,
	deps: []*Polyfill{  Window,  Console,  },
	versions: LastVersionBeforeSupport{},
}

var Console_warn = &Polyfill{
	code: Console_warn_Code,
	deps: []*Polyfill{  Window,  Console,  Console_log,  },
	versions: LastVersionBeforeSupport{},
}

var DevicePixelRatio = &Polyfill{
	code: DevicePixelRatio_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Document = &Polyfill{
	code: Document_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Document_currentScript = &Polyfill{
	code: Document_currentScript_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Document_getElementsByClassName = &Polyfill{
	code: Document_getElementsByClassName_Code,
	deps: []*Polyfill{  Document_querySelector,  },
	versions: LastVersionBeforeSupport{},
}

var Document_head = &Polyfill{
	code: Document_head_Code,
	deps: []*Polyfill{  Document,  },
	versions: LastVersionBeforeSupport{},
}

var Document_querySelector = &Polyfill{
	code: Document_querySelector_Code,
	deps: []*Polyfill{  Element,  Document,  },
	versions: LastVersionBeforeSupport{},
}

var Document_visibilityState = &Polyfill{
	code: Document_visibilityState_Code,
	deps: []*Polyfill{  CustomEvent,  },
	versions: LastVersionBeforeSupport{},
}

var Fetch = &Polyfill{
	code: Fetch_Code,
	deps: []*Polyfill{  Array_prototype_forEach,  Object_getOwnPropertyNames,  Promise,  XMLHttpRequest,  Symbol_iterator,  },
	versions: LastVersionBeforeSupport{},
}

var GetComputedStyle = &Polyfill{
	code: GetComputedStyle_Code,
	deps: []*Polyfill{  Window,  },
	versions: LastVersionBeforeSupport{},
}

var GlobalThis = &Polyfill{
	code: GlobalThis_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var LocalStorage = &Polyfill{
	code: LocalStorage_Code,
	deps: []*Polyfill{  Array_prototype_forEach,  Window,  },
	versions: LastVersionBeforeSupport{},
}

var Location_origin = &Polyfill{
	code: Location_origin_Code,
	deps: []*Polyfill{  Object_defineProperties,  },
	versions: LastVersionBeforeSupport{},
}

var MatchMedia = &Polyfill{
	code: MatchMedia_Code,
	deps: []*Polyfill{  Event,  },
	versions: LastVersionBeforeSupport{},
}

var Navigator_geolocation = &Polyfill{
	code: Navigator_geolocation_Code,
	deps: []*Polyfill{  Document_head,  },
	versions: LastVersionBeforeSupport{},
}

var Navigator_sendBeacon = &Polyfill{
	code: Navigator_sendBeacon_Code,
	deps: []*Polyfill{  XMLHttpRequest,  },
	versions: LastVersionBeforeSupport{},
}

var Performance_now = &Polyfill{
	code: Performance_now_Code,
	deps: []*Polyfill{  Date_now,  },
	versions: LastVersionBeforeSupport{},
}

var QueueMicrotask = &Polyfill{
	code: QueueMicrotask_Code,
	deps: []*Polyfill{  Promise,  Event,  },
	versions: LastVersionBeforeSupport{},
}

var RequestAnimationFrame = &Polyfill{
	code: RequestAnimationFrame_Code,
	deps: []*Polyfill{  Date_now,  },
	versions: LastVersionBeforeSupport{},
}

var RequestIdleCallback = &Polyfill{
	code: RequestIdleCallback_Code,
	deps: []*Polyfill{  Object_defineProperty,  Array_prototype_filter,  Array_prototype_includes,  RequestAnimationFrame,  Performance_now,  Event,  },
	versions: LastVersionBeforeSupport{},
}

var Screen_orientation = &Polyfill{
	code: Screen_orientation_Code,
	deps: []*Polyfill{  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var SetImmediate = &Polyfill{
	code: SetImmediate_Code,
	deps: []*Polyfill{  },
	versions: LastVersionBeforeSupport{},
}

var Smoothscroll = &Polyfill{
	code: Smoothscroll_Code,
	deps: []*Polyfill{  RequestAnimationFrame,  Function_prototype_bind,  GetComputedStyle,  Object_defineProperty,  },
	versions: LastVersionBeforeSupport{},
}

var __html5_elements = &Polyfill{
	code: __html5_elements_Code,
	deps: []*Polyfill{  Document,  },
	versions: LastVersionBeforeSupport{},
}

var __viewport = &Polyfill{
	code: __viewport_Code,
	deps: []*Polyfill{  Object_defineProperties,  },
	versions: LastVersionBeforeSupport{},
}


var polyfillsByName = map[string]*Polyfill{

	"AbortController": AbortController,
	"Array.from": Array_from,
	"Array.isArray": Array_isArray,
	"Array.of": Array_of,
	"Array.prototype.@@iterator": Array_prototype_ααiterator,
	"Array.prototype.copyWithin": Array_prototype_copyWithin,
	"Array.prototype.entries": Array_prototype_entries,
	"Array.prototype.every": Array_prototype_every,
	"Array.prototype.fill": Array_prototype_fill,
	"Array.prototype.filter": Array_prototype_filter,
	"Array.prototype.find": Array_prototype_find,
	"Array.prototype.findIndex": Array_prototype_findIndex,
	"Array.prototype.flat": Array_prototype_flat,
	"Array.prototype.flatMap": Array_prototype_flatMap,
	"Array.prototype.forEach": Array_prototype_forEach,
	"Array.prototype.includes": Array_prototype_includes,
	"Array.prototype.indexOf": Array_prototype_indexOf,
	"Array.prototype.keys": Array_prototype_keys,
	"Array.prototype.lastIndexOf": Array_prototype_lastIndexOf,
	"Array.prototype.map": Array_prototype_map,
	"Array.prototype.reduce": Array_prototype_reduce,
	"Array.prototype.reduceRight": Array_prototype_reduceRight,
	"Array.prototype.some": Array_prototype_some,
	"Array.prototype.sort": Array_prototype_sort,
	"Array.prototype.values": Array_prototype_values,
	"AudioContext": AudioContext,
	"Blob": Blob,
	"CSS.supports": CSS_supports,
	"CustomEvent": CustomEvent,
	"DOMRect": DOMRect,
	"DOMTokenList": DOMTokenList,
	"DOMTokenList.prototype.@@iterator": DOMTokenList_prototype_ααiterator,
	"DOMTokenList.prototype.forEach": DOMTokenList_prototype_forEach,
	"DOMTokenList.prototype.replace": DOMTokenList_prototype_replace,
	"Date.now": Date_now,
	"Date.prototype.toISOString": Date_prototype_toISOString,
	"DocumentFragment": DocumentFragment,
	"DocumentFragment.prototype.append": DocumentFragment_prototype_append,
	"DocumentFragment.prototype.prepend": DocumentFragment_prototype_prepend,
	"Element": Element,
	"Element.prototype.after": Element_prototype_after,
	"Element.prototype.append": Element_prototype_append,
	"Element.prototype.before": Element_prototype_before,
	"Element.prototype.classList": Element_prototype_classList,
	"Element.prototype.cloneNode": Element_prototype_cloneNode,
	"Element.prototype.closest": Element_prototype_closest,
	"Element.prototype.dataset": Element_prototype_dataset,
	"Element.prototype.inert": Element_prototype_inert,
	"Element.prototype.matches": Element_prototype_matches,
	"Element.prototype.nextElementSibling": Element_prototype_nextElementSibling,
	"Element.prototype.placeholder": Element_prototype_placeholder,
	"Element.prototype.prepend": Element_prototype_prepend,
	"Element.prototype.previousElementSibling": Element_prototype_previousElementSibling,
	"Element.prototype.remove": Element_prototype_remove,
	"Element.prototype.replaceWith": Element_prototype_replaceWith,
	"Element.prototype.toggleAttribute": Element_prototype_toggleAttribute,
	"Event": Event,
	"Event.focusin": Event_focusin,
	"Event.hashchange": Event_hashchange,
	"EventSource": EventSource,
	"Function.prototype.bind": Function_prototype_bind,
	"Function.prototype.name": Function_prototype_name,
	"HTMLCanvasElement.prototype.toBlob": HTMLCanvasElement_prototype_toBlob,
	"HTMLDocument": HTMLDocument,
	"HTMLPictureElement": HTMLPictureElement,
	"HTMLTemplateElement": HTMLTemplateElement,
	"IntersectionObserver": IntersectionObserver,
	"IntersectionObserverEntry": IntersectionObserverEntry,
	"Intl.DateTimeFormat": Intl_DateTimeFormat,
	"Intl.DisplayNames": Intl_DisplayNames,
	"Intl.ListFormat": Intl_ListFormat,
	"Intl.Locale": Intl_Locale,
	"Intl.NumberFormat": Intl_NumberFormat,
	"Intl.PluralRules": Intl_PluralRules,
	"Intl.RelativeTimeFormat": Intl_RelativeTimeFormat,
	"Intl.getCanonicalLocales": Intl_getCanonicalLocales,
	"JSON": JSON,
	"Map": Map,
	"Math.acosh": Math_acosh,
	"Math.asinh": Math_asinh,
	"Math.atanh": Math_atanh,
	"Math.cbrt": Math_cbrt,
	"Math.clz32": Math_clz32,
	"Math.cosh": Math_cosh,
	"Math.expm1": Math_expm1,
	"Math.fround": Math_fround,
	"Math.hypot": Math_hypot,
	"Math.imul": Math_imul,
	"Math.log10": Math_log10,
	"Math.log1p": Math_log1p,
	"Math.log2": Math_log2,
	"Math.sign": Math_sign,
	"Math.sinh": Math_sinh,
	"Math.tanh": Math_tanh,
	"Math.trunc": Math_trunc,
	"MutationObserver": MutationObserver,
	"Node.prototype.contains": Node_prototype_contains,
	"NodeList.prototype.@@iterator": NodeList_prototype_ααiterator,
	"NodeList.prototype.forEach": NodeList_prototype_forEach,
	"Number.Epsilon": Number_Epsilon,
	"Number.MAX_SAFE_INTEGER": Number_MAX_SAFE_INTEGER,
	"Number.MIN_SAFE_INTEGER": Number_MIN_SAFE_INTEGER,
	"Number.isFinite": Number_isFinite,
	"Number.isInteger": Number_isInteger,
	"Number.isNaN": Number_isNaN,
	"Number.isSafeInteger": Number_isSafeInteger,
	"Number.parseFloat": Number_parseFloat,
	"Number.parseInt": Number_parseInt,
	"Object.assign": Object_assign,
	"Object.create": Object_create,
	"Object.defineProperties": Object_defineProperties,
	"Object.defineProperty": Object_defineProperty,
	"Object.entries": Object_entries,
	"Object.freeze": Object_freeze,
	"Object.fromEntries": Object_fromEntries,
	"Object.getOwnPropertyDescriptor": Object_getOwnPropertyDescriptor,
	"Object.getOwnPropertyDescriptors": Object_getOwnPropertyDescriptors,
	"Object.getOwnPropertyNames": Object_getOwnPropertyNames,
	"Object.getPrototypeOf": Object_getPrototypeOf,
	"Object.is": Object_is,
	"Object.isExtensible": Object_isExtensible,
	"Object.isFrozen": Object_isFrozen,
	"Object.isSealed": Object_isSealed,
	"Object.keys": Object_keys,
	"Object.preventExtensions": Object_preventExtensions,
	"Object.seal": Object_seal,
	"Object.setPrototypeOf": Object_setPrototypeOf,
	"Object.values": Object_values,
	"Promise": Promise,
	"Promise.prototype.finally": Promise_prototype_finally,
	"Reflect": Reflect,
	"Reflect.apply": Reflect_apply,
	"Reflect.construct": Reflect_construct,
	"Reflect.defineProperty": Reflect_defineProperty,
	"Reflect.deleteProperty": Reflect_deleteProperty,
	"Reflect.get": Reflect_get,
	"Reflect.getOwnPropertyDescriptor": Reflect_getOwnPropertyDescriptor,
	"Reflect.getPrototypeOf": Reflect_getPrototypeOf,
	"Reflect.has": Reflect_has,
	"Reflect.isExtensible": Reflect_isExtensible,
	"Reflect.ownKeys": Reflect_ownKeys,
	"Reflect.preventExtensions": Reflect_preventExtensions,
	"Reflect.set": Reflect_set,
	"Reflect.setPrototypeOf": Reflect_setPrototypeOf,
	"RegExp.prototype.flags": RegExp_prototype_flags,
	"ResizeObserver": ResizeObserver,
	"Set": Set,
	"String.fromCodePoint": String_fromCodePoint,
	"String.prototype.@@iterator": String_prototype_ααiterator,
	"String.prototype.anchor": String_prototype_anchor,
	"String.prototype.big": String_prototype_big,
	"String.prototype.blink": String_prototype_blink,
	"String.prototype.bold": String_prototype_bold,
	"String.prototype.codePointAt": String_prototype_codePointAt,
	"String.prototype.endsWith": String_prototype_endsWith,
	"String.prototype.fixed": String_prototype_fixed,
	"String.prototype.fontcolor": String_prototype_fontcolor,
	"String.prototype.fontsize": String_prototype_fontsize,
	"String.prototype.includes": String_prototype_includes,
	"String.prototype.italics": String_prototype_italics,
	"String.prototype.link": String_prototype_link,
	"String.prototype.normalize": String_prototype_normalize,
	"String.prototype.padEnd": String_prototype_padEnd,
	"String.prototype.padStart": String_prototype_padStart,
	"String.prototype.repeat": String_prototype_repeat,
	"String.prototype.replaceAll": String_prototype_replaceAll,
	"String.prototype.small": String_prototype_small,
	"String.prototype.startsWith": String_prototype_startsWith,
	"String.prototype.strike": String_prototype_strike,
	"String.prototype.sub": String_prototype_sub,
	"String.prototype.sup": String_prototype_sup,
	"String.prototype.trim": String_prototype_trim,
	"String.prototype.trimEnd": String_prototype_trimEnd,
	"String.prototype.trimStart": String_prototype_trimStart,
	"String.raw": String_raw,
	"Symbol": Symbol,
	"Symbol.asyncIterator": Symbol_asyncIterator,
	"Symbol.hasInstance": Symbol_hasInstance,
	"Symbol.isConcatSpreadable": Symbol_isConcatSpreadable,
	"Symbol.iterator": Symbol_iterator,
	"Symbol.match": Symbol_match,
	"Symbol.prototype.description": Symbol_prototype_description,
	"Symbol.replace": Symbol_replace,
	"Symbol.search": Symbol_search,
	"Symbol.species": Symbol_species,
	"Symbol.split": Symbol_split,
	"Symbol.toPrimitive": Symbol_toPrimitive,
	"Symbol.toStringTag": Symbol_toStringTag,
	"Symbol.unscopables": Symbol_unscopables,
	"TextEncoder": TextEncoder,
	"URL": URL,
	"UserTiming": UserTiming,
	"WeakMap": WeakMap,
	"WeakSet": WeakSet,
	"WebAnimations": WebAnimations,
	"Window": Window,
	"XMLHttpRequest": XMLHttpRequest,
	"_ArrayIterator": _ArrayIterator,
	"_DOMTokenList": _DOMTokenList,
	"_ESAbstract.AddEntriesFromIterable": _ESAbstract_AddEntriesFromIterable,
	"_ESAbstract.ArrayCreate": _ESAbstract_ArrayCreate,
	"_ESAbstract.ArraySpeciesCreate": _ESAbstract_ArraySpeciesCreate,
	"_ESAbstract.Call": _ESAbstract_Call,
	"_ESAbstract.CanonicalNumericIndexString": _ESAbstract_CanonicalNumericIndexString,
	"_ESAbstract.Construct": _ESAbstract_Construct,
	"_ESAbstract.CreateDataProperty": _ESAbstract_CreateDataProperty,
	"_ESAbstract.CreateDataPropertyOrThrow": _ESAbstract_CreateDataPropertyOrThrow,
	"_ESAbstract.CreateHTML": _ESAbstract_CreateHTML,
	"_ESAbstract.CreateIterResultObject": _ESAbstract_CreateIterResultObject,
	"_ESAbstract.CreateListFromArrayLike": _ESAbstract_CreateListFromArrayLike,
	"_ESAbstract.CreateMethodProperty": _ESAbstract_CreateMethodProperty,
	"_ESAbstract.EnumerableOwnProperties": _ESAbstract_EnumerableOwnProperties,
	"_ESAbstract.FlattenIntoArray": _ESAbstract_FlattenIntoArray,
	"_ESAbstract.Get": _ESAbstract_Get,
	"_ESAbstract.GetIterator": _ESAbstract_GetIterator,
	"_ESAbstract.GetMethod": _ESAbstract_GetMethod,
	"_ESAbstract.GetPrototypeFromConstructor": _ESAbstract_GetPrototypeFromConstructor,
	"_ESAbstract.GetSubstitution": _ESAbstract_GetSubstitution,
	"_ESAbstract.GetV": _ESAbstract_GetV,
	"_ESAbstract.HasOwnProperty": _ESAbstract_HasOwnProperty,
	"_ESAbstract.HasProperty": _ESAbstract_HasProperty,
	"_ESAbstract.Invoke": _ESAbstract_Invoke,
	"_ESAbstract.IsArray": _ESAbstract_IsArray,
	"_ESAbstract.IsCallable": _ESAbstract_IsCallable,
	"_ESAbstract.IsConstructor": _ESAbstract_IsConstructor,
	"_ESAbstract.IsInteger": _ESAbstract_IsInteger,
	"_ESAbstract.IsPropertyKey": _ESAbstract_IsPropertyKey,
	"_ESAbstract.IsRegExp": _ESAbstract_IsRegExp,
	"_ESAbstract.IteratorClose": _ESAbstract_IteratorClose,
	"_ESAbstract.IteratorComplete": _ESAbstract_IteratorComplete,
	"_ESAbstract.IteratorNext": _ESAbstract_IteratorNext,
	"_ESAbstract.IteratorStep": _ESAbstract_IteratorStep,
	"_ESAbstract.IteratorValue": _ESAbstract_IteratorValue,
	"_ESAbstract.OrdinaryCreateFromConstructor": _ESAbstract_OrdinaryCreateFromConstructor,
	"_ESAbstract.OrdinaryToPrimitive": _ESAbstract_OrdinaryToPrimitive,
	"_ESAbstract.RequireObjectCoercible": _ESAbstract_RequireObjectCoercible,
	"_ESAbstract.SameValue": _ESAbstract_SameValue,
	"_ESAbstract.SameValueNonNumber": _ESAbstract_SameValueNonNumber,
	"_ESAbstract.SameValueZero": _ESAbstract_SameValueZero,
	"_ESAbstract.SpeciesConstructor": _ESAbstract_SpeciesConstructor,
	"_ESAbstract.StringIndexOf": _ESAbstract_StringIndexOf,
	"_ESAbstract.ToBoolean": _ESAbstract_ToBoolean,
	"_ESAbstract.ToIndex": _ESAbstract_ToIndex,
	"_ESAbstract.ToInt16": _ESAbstract_ToInt16,
	"_ESAbstract.ToInt32": _ESAbstract_ToInt32,
	"_ESAbstract.ToInt8": _ESAbstract_ToInt8,
	"_ESAbstract.ToInteger": _ESAbstract_ToInteger,
	"_ESAbstract.ToLength": _ESAbstract_ToLength,
	"_ESAbstract.ToNumber": _ESAbstract_ToNumber,
	"_ESAbstract.ToObject": _ESAbstract_ToObject,
	"_ESAbstract.ToPrimitive": _ESAbstract_ToPrimitive,
	"_ESAbstract.ToPropertyDescriptor": _ESAbstract_ToPropertyDescriptor,
	"_ESAbstract.ToPropertyKey": _ESAbstract_ToPropertyKey,
	"_ESAbstract.ToString": _ESAbstract_ToString,
	"_ESAbstract.ToUint16": _ESAbstract_ToUint16,
	"_ESAbstract.ToUint32": _ESAbstract_ToUint32,
	"_ESAbstract.ToUint8": _ESAbstract_ToUint8,
	"_ESAbstract.ToUint8Clamp": _ESAbstract_ToUint8Clamp,
	"_ESAbstract.TrimString": _ESAbstract_TrimString,
	"_ESAbstract.Type": _ESAbstract_Type,
	"_ESAbstract.UTF16Decode": _ESAbstract_UTF16Decode,
	"_ESAbstract.UTF16Encoding": _ESAbstract_UTF16Encoding,
	"_Iterator": _Iterator,
	"_StringIterator": _StringIterator,
	"_TypedArray": _TypedArray,
	"_mutation": _mutation,
	"atob": Atob,
	"console": Console,
	"console.assert": Console_assert,
	"console.clear": Console_clear,
	"console.count": Console_count,
	"console.debug": Console_debug,
	"console.dir": Console_dir,
	"console.dirxml": Console_dirxml,
	"console.error": Console_error,
	"console.exception": Console_exception,
	"console.group": Console_group,
	"console.groupCollapsed": Console_groupCollapsed,
	"console.groupEnd": Console_groupEnd,
	"console.info": Console_info,
	"console.log": Console_log,
	"console.markTimeline": Console_markTimeline,
	"console.profile": Console_profile,
	"console.profileEnd": Console_profileEnd,
	"console.profiles": Console_profiles,
	"console.table": Console_table,
	"console.time": Console_time,
	"console.timeEnd": Console_timeEnd,
	"console.timeStamp": Console_timeStamp,
	"console.timeline": Console_timeline,
	"console.timelineEnd": Console_timelineEnd,
	"console.trace": Console_trace,
	"console.warn": Console_warn,
	"devicePixelRatio": DevicePixelRatio,
	"document": Document,
	"document.currentScript": Document_currentScript,
	"document.getElementsByClassName": Document_getElementsByClassName,
	"document.head": Document_head,
	"document.querySelector": Document_querySelector,
	"document.visibilityState": Document_visibilityState,
	"fetch": Fetch,
	"getComputedStyle": GetComputedStyle,
	"globalThis": GlobalThis,
	"localStorage": LocalStorage,
	"location.origin": Location_origin,
	"matchMedia": MatchMedia,
	"navigator.geolocation": Navigator_geolocation,
	"navigator.sendBeacon": Navigator_sendBeacon,
	"performance.now": Performance_now,
	"queueMicrotask": QueueMicrotask,
	"requestAnimationFrame": RequestAnimationFrame,
	"requestIdleCallback": RequestIdleCallback,
	"screen.orientation": Screen_orientation,
	"setImmediate": SetImmediate,
	"smoothscroll": Smoothscroll,
	"~html5-elements": __html5_elements,
	"~viewport": __viewport,
}
