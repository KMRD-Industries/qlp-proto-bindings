// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: test.proto
// Protobuf C++ Version: 5.27.0-dev

#ifndef GOOGLE_PROTOBUF_INCLUDED_test_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_test_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>
#include <utility>

#include "google/protobuf/runtime_version.h"
#if PROTOBUF_VERSION != 5027000
#error "Protobuf C++ gencode is built with an incompatible version of"
#error "Protobuf C++ headers/runtime. See"
#error "https://protobuf.dev/support/cross-version-runtime-guarantee/#cpp"
#endif
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/generated_enum_reflection.h"
#include "google/protobuf/unknown_field_set.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_test_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_test_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_test_2eproto;
namespace comm {
class Data;
struct DataDefaultTypeInternal;
extern DataDefaultTypeInternal _Data_default_instance_;
}  // namespace comm
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace comm {
enum State : int {
  STATE_UNSPECIFIED = 0,
  STATE_ALIVE = 1,
  STATE_DEAD = 2,
  State_INT_MIN_SENTINEL_DO_NOT_USE_ =
      std::numeric_limits<::int32_t>::min(),
  State_INT_MAX_SENTINEL_DO_NOT_USE_ =
      std::numeric_limits<::int32_t>::max(),
};

bool State_IsValid(int value);
extern const uint32_t State_internal_data_[];
constexpr State State_MIN = static_cast<State>(0);
constexpr State State_MAX = static_cast<State>(2);
constexpr int State_ARRAYSIZE = 2 + 1;
const ::google::protobuf::EnumDescriptor*
State_descriptor();
template <typename T>
const std::string& State_Name(T value) {
  static_assert(std::is_same<T, State>::value ||
                    std::is_integral<T>::value,
                "Incorrect type passed to State_Name().");
  return State_Name(static_cast<State>(value));
}
template <>
inline const std::string& State_Name(State value) {
  return ::google::protobuf::internal::NameOfDenseEnum<State_descriptor,
                                                 0, 2>(
      static_cast<int>(value));
}
inline bool State_Parse(absl::string_view name, State* value) {
  return ::google::protobuf::internal::ParseNamedEnum<State>(
      State_descriptor(), name, value);
}

// ===================================================================


// -------------------------------------------------------------------

class Data final : public ::google::protobuf::Message
/* @@protoc_insertion_point(class_definition:comm.Data) */ {
 public:
  inline Data() : Data(nullptr) {}
  ~Data() override;
  template <typename = void>
  explicit PROTOBUF_CONSTEXPR Data(
      ::google::protobuf::internal::ConstantInitialized);

  inline Data(const Data& from) : Data(nullptr, from) {}
  inline Data(Data&& from) noexcept
      : Data(nullptr, std::move(from)) {}
  inline Data& operator=(const Data& from) {
    CopyFrom(from);
    return *this;
  }
  inline Data& operator=(Data&& from) noexcept {
    if (this == &from) return *this;
    if (GetArena() == from.GetArena()
#ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetArena() != nullptr
#endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields()
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Data& default_instance() {
    return *internal_default_instance();
  }
  static inline const Data* internal_default_instance() {
    return reinterpret_cast<const Data*>(
        &_Data_default_instance_);
  }
  static constexpr int kIndexInFileMessages = 0;
  friend void swap(Data& a, Data& b) { a.Swap(&b); }
  inline void Swap(Data* other) {
    if (other == this) return;
#ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() != nullptr && GetArena() == other->GetArena()) {
#else   // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() == other->GetArena()) {
#endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Data* other) {
    if (other == this) return;
    ABSL_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Data* New(::google::protobuf::Arena* arena = nullptr) const final {
    return ::google::protobuf::Message::DefaultConstruct<Data>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const Data& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom(const Data& from) { Data::MergeImpl(*this, from); }

  private:
  static void MergeImpl(
      ::google::protobuf::MessageLite& to_msg,
      const ::google::protobuf::MessageLite& from_msg);

  public:
  ABSL_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  ::size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::google::protobuf::internal::ParseContext* ctx) final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target,
      ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void InternalSwap(Data* other);
 private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() { return "comm.Data"; }

 protected:
  explicit Data(::google::protobuf::Arena* arena);
  Data(::google::protobuf::Arena* arena, const Data& from);
  Data(::google::protobuf::Arena* arena, Data&& from) noexcept
      : Data(arena) {
    *this = ::std::move(from);
  }
  const ::google::protobuf::Message::ClassData* GetClassData() const final;

 public:
  ::google::protobuf::Metadata GetMetadata() const;
  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------
  enum : int {
    kMsgListFieldNumber = 4,
    kAFieldNumber = 2,
    kIFieldNumber = 1,
    kStateFieldNumber = 3,
  };
  // repeated string msg_list = 4;
  int msg_list_size() const;
  private:
  int _internal_msg_list_size() const;

  public:
  void clear_msg_list() ;
  const std::string& msg_list(int index) const;
  std::string* mutable_msg_list(int index);
  void set_msg_list(int index, const std::string& value);
  void set_msg_list(int index, std::string&& value);
  void set_msg_list(int index, const char* value);
  void set_msg_list(int index, const char* value, std::size_t size);
  void set_msg_list(int index, absl::string_view value);
  std::string* add_msg_list();
  void add_msg_list(const std::string& value);
  void add_msg_list(std::string&& value);
  void add_msg_list(const char* value);
  void add_msg_list(const char* value, std::size_t size);
  void add_msg_list(absl::string_view value);
  const ::google::protobuf::RepeatedPtrField<std::string>& msg_list() const;
  ::google::protobuf::RepeatedPtrField<std::string>* mutable_msg_list();

  private:
  const ::google::protobuf::RepeatedPtrField<std::string>& _internal_msg_list() const;
  ::google::protobuf::RepeatedPtrField<std::string>* _internal_mutable_msg_list();

  public:
  // optional string a = 2;
  bool has_a() const;
  void clear_a() ;
  const std::string& a() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_a(Arg_&& arg, Args_... args);
  std::string* mutable_a();
  PROTOBUF_NODISCARD std::string* release_a();
  void set_allocated_a(std::string* value);

  private:
  const std::string& _internal_a() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_a(
      const std::string& value);
  std::string* _internal_mutable_a();

  public:
  // optional int32 i = 1;
  bool has_i() const;
  void clear_i() ;
  ::int32_t i() const;
  void set_i(::int32_t value);

  private:
  ::int32_t _internal_i() const;
  void _internal_set_i(::int32_t value);

  public:
  // .comm.State state = 3;
  void clear_state() ;
  ::comm::State state() const;
  void set_state(::comm::State value);

  private:
  ::comm::State _internal_state() const;
  void _internal_set_state(::comm::State value);

  public:
  // @@protoc_insertion_point(class_scope:comm.Data)
 private:
  class _Internal;
  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<
      2, 4, 0,
      27, 2>
      _table_;

  static constexpr const void* _raw_default_instance_ =
      &_Data_default_instance_;

  friend class ::google::protobuf::MessageLite;
  friend class ::google::protobuf::Arena;
  template <typename T>
  friend class ::google::protobuf::Arena::InternalHelper;
  using InternalArenaConstructable_ = void;
  using DestructorSkippable_ = void;
  struct Impl_ {
    inline explicit constexpr Impl_(
        ::google::protobuf::internal::ConstantInitialized) noexcept;
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena);
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena, const Impl_& from);
    ::google::protobuf::internal::HasBits<1> _has_bits_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    ::google::protobuf::RepeatedPtrField<std::string> msg_list_;
    ::google::protobuf::internal::ArenaStringPtr a_;
    ::int32_t i_;
    int state_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_test_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// Data

// optional int32 i = 1;
inline bool Data::has_i() const {
  bool value = (_impl_._has_bits_[0] & 0x00000002u) != 0;
  return value;
}
inline void Data::clear_i() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.i_ = 0;
  _impl_._has_bits_[0] &= ~0x00000002u;
}
inline ::int32_t Data::i() const {
  // @@protoc_insertion_point(field_get:comm.Data.i)
  return _internal_i();
}
inline void Data::set_i(::int32_t value) {
  _internal_set_i(value);
  _impl_._has_bits_[0] |= 0x00000002u;
  // @@protoc_insertion_point(field_set:comm.Data.i)
}
inline ::int32_t Data::_internal_i() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.i_;
}
inline void Data::_internal_set_i(::int32_t value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.i_ = value;
}

// optional string a = 2;
inline bool Data::has_a() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  return value;
}
inline void Data::clear_a() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.a_.ClearToEmpty();
  _impl_._has_bits_[0] &= ~0x00000001u;
}
inline const std::string& Data::a() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:comm.Data.a)
  return _internal_a();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void Data::set_a(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_._has_bits_[0] |= 0x00000001u;
  _impl_.a_.Set(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:comm.Data.a)
}
inline std::string* Data::mutable_a() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_a();
  // @@protoc_insertion_point(field_mutable:comm.Data.a)
  return _s;
}
inline const std::string& Data::_internal_a() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.a_.Get();
}
inline void Data::_internal_set_a(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_._has_bits_[0] |= 0x00000001u;
  _impl_.a_.Set(value, GetArena());
}
inline std::string* Data::_internal_mutable_a() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_._has_bits_[0] |= 0x00000001u;
  return _impl_.a_.Mutable( GetArena());
}
inline std::string* Data::release_a() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:comm.Data.a)
  if ((_impl_._has_bits_[0] & 0x00000001u) == 0) {
    return nullptr;
  }
  _impl_._has_bits_[0] &= ~0x00000001u;
  auto* released = _impl_.a_.Release();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  _impl_.a_.Set("", GetArena());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  return released;
}
inline void Data::set_allocated_a(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  if (value != nullptr) {
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }
  _impl_.a_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.a_.IsDefault()) {
          _impl_.a_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:comm.Data.a)
}

// .comm.State state = 3;
inline void Data::clear_state() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.state_ = 0;
}
inline ::comm::State Data::state() const {
  // @@protoc_insertion_point(field_get:comm.Data.state)
  return _internal_state();
}
inline void Data::set_state(::comm::State value) {
  _internal_set_state(value);
  // @@protoc_insertion_point(field_set:comm.Data.state)
}
inline ::comm::State Data::_internal_state() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return static_cast<::comm::State>(_impl_.state_);
}
inline void Data::_internal_set_state(::comm::State value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.state_ = value;
}

// repeated string msg_list = 4;
inline int Data::_internal_msg_list_size() const {
  return _internal_msg_list().size();
}
inline int Data::msg_list_size() const {
  return _internal_msg_list_size();
}
inline void Data::clear_msg_list() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.msg_list_.Clear();
}
inline std::string* Data::add_msg_list()
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  std::string* _s = _internal_mutable_msg_list()->Add();
  // @@protoc_insertion_point(field_add_mutable:comm.Data.msg_list)
  return _s;
}
inline const std::string& Data::msg_list(int index) const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:comm.Data.msg_list)
  return _internal_msg_list().Get(index);
}
inline std::string* Data::mutable_msg_list(int index)
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_mutable:comm.Data.msg_list)
  return _internal_mutable_msg_list()->Mutable(index);
}
inline void Data::set_msg_list(int index, const std::string& value) {
  _internal_mutable_msg_list()->Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set:comm.Data.msg_list)
}
inline void Data::set_msg_list(int index, std::string&& value) {
  _internal_mutable_msg_list()->Mutable(index)->assign(std::move(value));
  // @@protoc_insertion_point(field_set:comm.Data.msg_list)
}
inline void Data::set_msg_list(int index, const char* value) {
  ABSL_DCHECK(value != nullptr);
  _internal_mutable_msg_list()->Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.Data.msg_list)
}
inline void Data::set_msg_list(int index, const char* value,
                              std::size_t size) {
  _internal_mutable_msg_list()->Mutable(index)->assign(
      reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.Data.msg_list)
}
inline void Data::set_msg_list(int index, absl::string_view value) {
  _internal_mutable_msg_list()->Mutable(index)->assign(
      value.data(), value.size());
  // @@protoc_insertion_point(field_set_string_piece:comm.Data.msg_list)
}
inline void Data::add_msg_list(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _internal_mutable_msg_list()->Add()->assign(value);
  // @@protoc_insertion_point(field_add:comm.Data.msg_list)
}
inline void Data::add_msg_list(std::string&& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _internal_mutable_msg_list()->Add(std::move(value));
  // @@protoc_insertion_point(field_add:comm.Data.msg_list)
}
inline void Data::add_msg_list(const char* value) {
  ABSL_DCHECK(value != nullptr);
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _internal_mutable_msg_list()->Add()->assign(value);
  // @@protoc_insertion_point(field_add_char:comm.Data.msg_list)
}
inline void Data::add_msg_list(const char* value, std::size_t size) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _internal_mutable_msg_list()->Add()->assign(
      reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_add_pointer:comm.Data.msg_list)
}
inline void Data::add_msg_list(absl::string_view value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _internal_mutable_msg_list()->Add()->assign(value.data(),
                                                     value.size());
  // @@protoc_insertion_point(field_add_string_piece:comm.Data.msg_list)
}
inline const ::google::protobuf::RepeatedPtrField<std::string>&
Data::msg_list() const ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_list:comm.Data.msg_list)
  return _internal_msg_list();
}
inline ::google::protobuf::RepeatedPtrField<std::string>*
Data::mutable_msg_list() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_mutable_list:comm.Data.msg_list)
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _internal_mutable_msg_list();
}
inline const ::google::protobuf::RepeatedPtrField<std::string>&
Data::_internal_msg_list() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.msg_list_;
}
inline ::google::protobuf::RepeatedPtrField<std::string>*
Data::_internal_mutable_msg_list() {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return &_impl_.msg_list_;
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace comm


namespace google {
namespace protobuf {

template <>
struct is_proto_enum<::comm::State> : std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor<::comm::State>() {
  return ::comm::State_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_test_2eproto_2epb_2eh